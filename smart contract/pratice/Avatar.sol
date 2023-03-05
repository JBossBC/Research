// SPDX-License-Identifier: MITu
pragma solidity 0.8.14;

// ReetrancyGuard
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./Bucket.sol";

contract Avatar is ReentrancyGuard, Bucket {
    //保本比例
    uint256 public constant PRINCIPAL_RATIO = 650000; // 65%
    //归属于静态奖励
    uint256 public constant INVEST_RATIO = 260000; // 26%
    //平台抽成
    uint256 public constant PLATFORM_RATIO = 20000; // 2%
    //动态奖励池
    uint256 public constant REFERRER_RATIO = 60000; // 6%
    //福利池(最后的百分之一)
    uint256 public constant INCENTIVE_RATIO = 10000; // 1%
    //价格精度
    uint256 public constant PRICE_PRECISION = 1e6;
    //默认投资回报率(会根据保本数量和目标数量进行调整,最小为10000-14*500)
    uint256 public constant DEFAULT_INVEST_RETURN_RATE = 10000; // 1%
    //使投资回报率增加
    uint256 public constant BOOST_INVEST_RETURN_RATE = 5000; // 0.5%
   //最大投资
    uint256 public constant MAX_INVEST = 1e21; // 1000
    //最小投资
    uint256 public constant MIN_INVEST = 1e20; // 100
    //时间单元(每一个时间单元会更新募集目标)
    uint256 public constant TIME_UNIT = 1 days;
    //默认目标金额
    uint256[6] public DEFAULT_TARGET_AMOUNTS = [13e22, 25e22, 35e22, 50e22, 75e22, 125e22];
    //动态投资所需要维护一个巨大的refer的库(具体用于_distributeReferrerReward函数)
    uint256 public constant MAX_SEARCH_DEPTH = 50;
    //分级激励
    uint256 public constant RANKED_INCENTIVE = 60;
    //平台地址(how to release the paymentsplitter contact address)
    address public platformAddress; // will be paymentsplitter contract address
    // 当前纪元 
    uint256[6] public currentEpochs;

    // ledge type => round epoch => address => position index => position info
    mapping(uint256 => mapping(address => PositionInfo[]))[6] public roundLedgers;
    //对应轮数的静态投资等级的信息
    mapping(uint256 => RoundInfo)[6] public roundInfos;
    //用户所在每轮的对应静态投资等级的信息,userRoundInfo[]数组下标代表纪元
    mapping(address => UserRoundInfo[])[6] public userRoundsInfos;
    //用户全局信息
    mapping(address => UserGlobalInfo) public userGlobalInfos;

    mapping(address => address[]) public children; // used for easily retrieve the referrer tree structure from front-end

    // temp admin
    address public tempAdmin;
    address public operator;
    bool public gamePaused;
    struct FundTarget {
        uint256 lastCheckTime;
        uint256 amount;
        uint256 achievedAmount;
    }

    struct UserGlobalInfo {
        // referrer chain to record the referrer relationship
        address referrer;
        // referrer rearward vault
        uint256 totalReferrerReward;
        //已经拿了的
        uint256 referrerRewardClaimed;
        // boost credit
        uint256 boostCredit;
        // sales record
        uint256 maxChildrenSales;
        //拉取投资的金额
        uint256 sales;
        //总共投资金额
        uint256 totalPositionAmount;
        uint256 reportedSales;
        uint8 salesLevel;
    }
    //盒子投资
    struct PositionInfo {
        uint256 amount;
        uint256 openTime;
        uint256 expiryTime;
        //投资回报比例
        uint256 investReturnRate;
        //回撤金额
        uint256 withdrawnAmount;
        //激励金额(可能和上级能获得多少金额相关)
        uint256 incentiveAmount;
        //投资回报金额
        uint256 investReturnAmount;
        //索引位置
        uint256 index;
        bool incentiveClaimable;
    }

    struct LinkedPosition {
        address user;
        uint256 userPositionIndex;
    }
     //每轮信息，不爆仓就不会结束
    struct RoundInfo {
        FundTarget fundTarget;
        //总共的静态投资金额
        uint256 totalPositionAmount; // total amount of all positions
        uint256 currentPrincipalAmount; // current principal amount
        //当前投资金额
        uint256 currentInvestAmount; // current invest amount
        //总共静态投资数量
        uint256 totalPositionCount; // total count of all positions
        //当前静态投资数量
        uint256 currentPositionCount; // total count of all open positions
        //当前激励金额
        uint256 currentIncentiveAmount; // current incentive amount
        //激励快照(所对应的最后六十个存的总金额)
        uint256 incentiveSnapshot; // check total position of last N positions
        //链表连接每个盒子(能够获得激励的盒子)
        uint256 head; // head of linked position for last N positions
        mapping(uint256 => LinkedPosition) linkedPositions; // used for incentive track
        //记录用户所在的最后进入的纪元
        mapping(address => uint256) ledgerRoundToUserRoundIndex; // this round index in userRoundsInfos
        bool stopLoss; // default false means the round is running
    }
    //记录用户对Avatar这款游戏的全部信息
    struct UserRoundInfo {
        //轮数
        uint256 epoch;
        //静态投资数量
        uint256 totalPositionAmount;
        //当前本金金额(怎么去算的)
        uint256 currentPrincipalAmount;
        //总共回撤的金额
        uint256 totalWithdrawnAmount;
        //总共
        uint256 totalIncentiveClaimedAmount;
        //已关闭的静态激励数量
        uint256 totalClosedPositionCount;
        //TODO 怎样去理解
        uint256 returnRateBoostedAmount;
    }
   //维护动态投资的全局状态
    struct ReferrerSearch {
        uint256 currentUserSales;
        uint256 currentReferrerSales;
        address currentReferrer;
        uint256 currentReferrerAmount;
        uint256 levelDiffAmount;
        uint256 leftLevelDiffAmount;
        uint256 levelDiffAmountPerLevel;
        uint256 levelSearchAmount;
        uint256 leftLevelSearchAmount;
        uint256 levelSearchAmountPerReferrer;
        uint256 levelSearchSales;
        uint256 currentReferrerMaxChildSales;
        uint256 currentUserTotalPosAmount;
        uint256 currentUserReportedSales;
        address currentUser;
        uint8 depth;
        uint8 levelSearchStep;
        uint8 currentLevelDiff;
        uint8 numLevelSearchCandidate;
        uint8 baseSalesLevel;
        uint8 currentReferrerLevel;
        //
        bool levelDiffDone;
        bool levelSearchDone;
        bool levelSalesDone;
    }
    //收益静态投资的参数结构体
    struct OpenPositionParams {
        uint256 principalAmount;
        uint256 investAmount;
        uint256 referrerAmount;
        uint256 incentiveAmount;
        uint256 investReturnRate;
    }

    event PositionOpened(
        address indexed user,
        uint256 indexed ledgeType,
        uint256 indexed epoch,
        uint256 positionIndex,
        uint256 amount
    );

    event PositionClosed(
        address indexed user,
        uint256 indexed ledgeType,
        uint256 indexed epoch,
        uint256 positionIndex,
        uint256 amount
    );

    event NewReferrer(address indexed user, address indexed referrer);
    event NewRound(uint256 indexed epoch, uint256 indexed ledgeType);
    event ReferrerRewardAdded(address indexed user, uint256 amount, uint256 indexed rewardType); // type 0 for levelDiff, 1 for levelSearch, 2 for levelSearch
    event ReferrerRewardClaimed(address indexed user, uint256 amount);
    event SalesLevelUpdated(address indexed user, uint8 level);
    event IncentiveClaimed(address indexed user, uint256 amount);

    modifier notContract() {
        require(msg.sender == tx.origin, "Contract not allowed");
        _;
    }

    /**
     * @param _platformAddress The address of the platform
     * @param _tempAdmin The address of the temp admin
     * @param _operator The address of the operator
     */
    constructor(
        address _platformAddress,
        address _tempAdmin,
        address _operator
    ) {
        require(
            _platformAddress != address(0) && _tempAdmin != address(0) && _operator != address(0),
            "Invalid address provided"
        );
        emit NewRound(0, 0);
        emit NewRound(0, 1);
        emit NewRound(0, 2);
        emit NewRound(0, 3);
        emit NewRound(0, 4);
        emit NewRound(0, 5);

        tempAdmin = _tempAdmin;
        operator = _operator;
        platformAddress = _platformAddress;
        gamePaused = true;
    }

    /**
     * @notice Set the game paused status
     * @param _paused: The game paused status
     */
    //启停游戏(需要operator权限),需要保证在游戏没有停止的时候tempAdmin需要存在
    function setPause(bool _paused) external {
        require(msg.sender == operator, "Only operator");
        // make sure the admin has dropped when game is unpaused
        //清除admin后才能开启游戏
        if (!_paused) {
            require(tempAdmin == address(0), "Temp admin not dropped");
        }
        gamePaused = _paused;
    }

    /**
     * @notice Transfer operator
     */
    //前提:消息调用者时operator并且入参地址不能缺省(函数功能:修改operator)
    function transferOperator(address _operator) external {
        require(msg.sender == operator, "Only operator");
        require(_operator != address(0), "Invalid address");
        operator = _operator;
    }

    /**
     * @notice Drop the temp admin privilege
     */
    //将tempAdmin下线(需要保证消息调用者是tempAdmin自己)
    function dropTempAdmin() external {
        require(msg.sender == tempAdmin, "Only admin");
        tempAdmin = address(0);
    }

    /**
     * @notice Batch set referrer information for users
     * @param users: The users to set
     * @param referrers: The referrers to set
     * @param salesLevels: The sales levels to set
     */
    //为用户批量设置引用者信息(//TODO 什么是引用者)
    //(条件:需要调用者拥有admin权限,并且入参的三个长度必须相等,同时前两个入参的每个元素必须不能为缺省地址)
    //过程:找到每个user的全局信息,然后对每个user的全局信息中的referer和saleslevel附上对应下标的referrers和saleslevels
    //最后将user放入children中对应的referres[i]中
    function batchSetReferrerInfo(
        address[] calldata users,
        address[] calldata referrers,
        uint8[] calldata salesLevels
    ) external {
        require(msg.sender == tempAdmin, "Only admin");
        require(users.length == referrers.length && users.length == salesLevels.length, "Invalid input");
        UserGlobalInfo storage userGlobalInfo;
        uint256 userLength = users.length;
        for (uint256 i = 0; i < userLength; ++i) {
            require(users[i] != address(0), "Invalid address provided");
            userGlobalInfo = userGlobalInfos[users[i]];
            require(userGlobalInfo.referrer == address(0), "Referrer already set");
            userGlobalInfo.referrer = referrers[i];
            userGlobalInfo.salesLevel = salesLevels[i];
            children[referrers[i]].push(users[i]);
        }
    }

    /**
     * @notice Set fixed stock distribution to specific ledger type
     * @param ledgerType: The ledger type to set
     * @param typeDays: The days to set
     * @param stock: The stock to set
     */
    //设置ledgerType的类型
    //前提:调用者有admin权限，设置的ledgerType在规定范围内,stock数组长度大于0,并且和typeDays的长度一样
    //TODO stock和typeDays代表的意义
    function setStock(
        uint256 ledgerType,
        uint8[] calldata typeDays,
        uint16[] calldata stock
    ) external {
        require(ledgerType > 0, "Invalid ledger type");
        require(ledgerType < 6, "Invalid ledger type");
        require(msg.sender == tempAdmin, "Only admin");
        require(stock.length > 0, "Invalid stock array");
        require(typeDays.length == stock.length, "Invalid params");

        _setStock(ledgerType, typeDays, stock);
    }

    /**
     * @notice Open a new position
     * @param ledgerType: The ledger type to open
     * @param targetEpoch: The target epoch to open
     * @param targetRate: The target ratio to open
     * @param referrer: The expected referrer
     * @param useBoost: Whether to use boost for this position
     */
    /**
     * __distributeReferrerReward()函数 1094行进行更新了全局状态
     */
    function openPosition(
        uint256 ledgerType,
        uint256 targetEpoch,
        uint256 targetRate,
        address referrer,
        bool useBoost
    ) external payable notContract nonReentrant {
        require(ledgerType < 6, "Invalid ledger type");
        require(targetEpoch == currentEpochs[ledgerType], "Invalid epoch");
        require(msg.value >= MIN_INVEST, "Too small");
        require(msg.value <= MAX_INVEST, "Too large");
        require(!gamePaused, "Paused");

        // load user global info
        UserGlobalInfo storage userGlobalInfo = userGlobalInfos[msg.sender];
        // load global round info
        RoundInfo storage roundInfo = roundInfos[ledgerType][targetEpoch];
        // placeholder for user round info
        UserRoundInfo storage userRoundInfo;

        // determine referrer
        //将购买的用户添加到referrer的下级中
        {
            address _referrer = userGlobalInfo.referrer;
            // if referrer is already set or msg.sender is the root user whose referrer is address(0)
            if (_referrer == address(0) && children[msg.sender].length == 0) {
                // if referrer is not set, set it and make sure it is a valid referrer
                require(referrer != address(0) && referrer != msg.sender, "Invalid referrer");
                // make sure referrer is registered already
                require(
                    userGlobalInfos[referrer].referrer != address(0) || children[referrer].length > 0,
                    "Invalid referrer"
                );

                // update storage
                userGlobalInfo.referrer = referrer;
                children[referrer].push(msg.sender);
                emit NewReferrer(msg.sender, referrer);
            }
        }

        // calculate each part of the amount
        OpenPositionParams memory params = OpenPositionParams({
            principalAmount: (msg.value * PRINCIPAL_RATIO) / PRICE_PRECISION,
            investAmount: (msg.value * INVEST_RATIO) / PRICE_PRECISION,
            referrerAmount: (msg.value * REFERRER_RATIO) / PRICE_PRECISION,
            incentiveAmount: (msg.value * INCENTIVE_RATIO) / PRICE_PRECISION,
            investReturnRate: _safeProcessFundTargetGetInvestReturnRate(roundInfo, ledgerType)
        });

        // check target ratio
        require(targetRate <= params.investReturnRate, "Invalid ratio");

        // update user's current ledger and current round info
        uint256 userRoundInfoLength = userRoundsInfos[ledgerType][msg.sender].length;
        if (
            userRoundInfoLength == 0 ||
            userRoundsInfos[ledgerType][msg.sender][userRoundInfoLength - 1].epoch < targetEpoch
        ) {
            // this is users first position in this round of this ledger type
            UserRoundInfo memory _userRoundInfo;
            _userRoundInfo = UserRoundInfo({
                epoch: targetEpoch,
                totalPositionAmount: 0,
                currentPrincipalAmount: 0,
                totalWithdrawnAmount: 0,
                totalIncentiveClaimedAmount: 0,
                totalClosedPositionCount: 0,
                returnRateBoostedAmount: 0
            });
            // push roundInfo to storage
            userRoundsInfos[ledgerType][msg.sender].push(_userRoundInfo);
            roundInfo.ledgerRoundToUserRoundIndex[msg.sender] = userRoundInfoLength;
            userRoundInfoLength += 1;
        }

        // fetch back the roundInfo from storage for further direct modification
        //添加当前纪元的用户信息
        userRoundInfo = userRoundsInfos[ledgerType][msg.sender][userRoundInfoLength - 1];
        userRoundInfo.totalPositionAmount += msg.value;
        userRoundInfo.currentPrincipalAmount += params.principalAmount;
       //增加收益的点(前提:boostcredit>=msg.value)//TODO boostcredit的初始化!!!
        if (useBoost) {
            uint256 boostCredit = userGlobalInfo.boostCredit;
            require(boostCredit >= msg.value, "Exceed boost credit");
            params.investReturnRate += BOOST_INVEST_RETURN_RATE; // + 0.5%
            userGlobalInfo.boostCredit -= msg.value;
        }

        // update ledger round info
        //修改本纪元对应静态投资等级的变量
        roundInfo.totalPositionAmount += msg.value;
        roundInfo.currentPrincipalAmount += params.principalAmount;
        roundInfo.currentInvestAmount += params.investAmount;
        roundInfo.currentPositionCount += 1;
        roundInfo.currentIncentiveAmount += params.incentiveAmount;
        roundInfo.incentiveSnapshot += msg.value;
        roundInfo.totalPositionCount += 1;
        //用户对应纪元的静态等级的总数赋予index
        uint256 userTotalPositionCount = roundLedgers[ledgerType][targetEpoch][msg.sender].length;
        // construct position info
        //将创建的静态投资盒子放入
        {
            uint256 expiryTime = block.timestamp;
            if (ledgerType == 0) {
                expiryTime += TIME_UNIT;
            } else {
                expiryTime += _pickDay(ledgerType, roundInfo.totalPositionCount) * TIME_UNIT;
            }

            PositionInfo memory positionInfo = PositionInfo({
                amount: msg.value,
                openTime: block.timestamp,
                expiryTime: expiryTime,
                investReturnRate: params.investReturnRate,
                withdrawnAmount: 0,     //***此时的incentiveAmount为 0  */
                incentiveAmount: 0, //***此时的incentiveAmount为 0  */
                investReturnAmount: 0, //***此时的incentiveAmount为 0  */
                index: userTotalPositionCount,
                incentiveClaimable: true
            });

            // push position info to round ledgers
            roundLedgers[ledgerType][targetEpoch][msg.sender].push(positionInfo);
        }

        // distribute referrer funds
        //当用户买盒子的时候，其上级立即获得奖励(来自推广池)
        _distributeReferrerReward(msg.sender, params.referrerAmount);
      //将创建盒子的创建者信息赋予linkedPosition
        {
            // ranked incentive track
            mapping(uint256 => LinkedPosition) storage linkedPositions = roundInfo.linkedPositions;

            // update the latest position (which is the current position) node
            LinkedPosition storage linkedPosition = linkedPositions[roundInfo.totalPositionCount - 1];
            linkedPosition.user = msg.sender;
            linkedPosition.userPositionIndex = userTotalPositionCount;

            // adjust head in order to keep track last N positions
            //超过60个盒子的话，盒子的激励奖励失效(滑动窗口)roundInfo.head指向能够获得激励奖励的头节点
            if (roundInfo.totalPositionCount - roundInfo.head > RANKED_INCENTIVE) {
                // fetch current head node
                LinkedPosition storage headLinkedPosition = linkedPositions[roundInfo.head];
                PositionInfo storage headPositionInfo = roundLedgers[ledgerType][targetEpoch][headLinkedPosition.user][
                    headLinkedPosition.userPositionIndex
                ];
                // previous head position now is not eligible for incentive
                headPositionInfo.incentiveClaimable = false;
                // subtract head position amount, because we only keep the last RANKED_INCENTIVE positions
                roundInfo.incentiveSnapshot -= headPositionInfo.amount;
                // shift head to next global position to keep track the last N positions
                roundInfo.head += 1;
            }
        }

        // do transfer to platform
        {
            //转账给平台响应的收益 2%个点
            (bool success, ) = platformAddress.call{
                value: msg.value -
                    params.principalAmount -
                    params.investAmount -
                    params.referrerAmount -
                    params.incentiveAmount
            }("");
            require(success, "Transfer failed.");
        }
        // emit event
        emit PositionOpened(msg.sender, ledgerType, targetEpoch, userTotalPositionCount, msg.value);
    }

    /**
     * @notice Close position
     * @param ledgerType: Ledger type
     * @param epoch: Epoch of the ledger
     * @param positionIndex: Position index of the user
     */
    //到期结算
    function closePosition(
        uint256 ledgerType,
        uint256 epoch,
        uint256 positionIndex
    ) external notContract nonReentrant {
        require(ledgerType < 6, "Invalid ledger type");
        require(epoch <= currentEpochs[ledgerType], "Invalid epoch");

        // check index is valid
        PositionInfo[] storage positionInfos = roundLedgers[ledgerType][epoch][msg.sender];
        require(positionIndex < positionInfos.length, "Invalid position index");

        // get position Info
        PositionInfo storage positionInfo = positionInfos[positionIndex];

        // get roundIno
        RoundInfo storage roundInfo = roundInfos[ledgerType][epoch];

        // user global info
        UserGlobalInfo storage userGlobalInfo = userGlobalInfos[msg.sender];

        _safeClosePosition(ledgerType, epoch, positionIndex, positionInfo, roundInfo, userGlobalInfo);
    }

    /**
     * @notice Close a batch of positions
     * @param ledgerType: Ledger type
     * @param epoch: Epoch of the ledger
     * @param positionIndexes: Position indexes of the user
     */
    function batchClosePositions(
        uint256 ledgerType,
        uint256 epoch,
        uint256[] calldata positionIndexes
    ) external nonReentrant {
        require(ledgerType < 6, "Invalid ledger type");
        require(epoch <= currentEpochs[ledgerType], "Invalid epoch");
        require(positionIndexes.length > 0, "Invalid position indexes");

        // check index is valid
        PositionInfo[] storage positionInfos = roundLedgers[ledgerType][epoch][msg.sender];

        // get roundIno
        RoundInfo storage roundInfo = roundInfos[ledgerType][epoch];

        // position info placeholder
        PositionInfo storage positionInfo;

        // user global info
        UserGlobalInfo storage userGlobalInfo = userGlobalInfos[msg.sender];

        uint256 positionIndexesLength = positionIndexes.length;
        uint256 positionInfosLength = positionInfos.length;
        for (uint256 i = 0; i < positionIndexesLength; ++i) {
            require(positionIndexes[i] < positionInfosLength, "Invalid position index");
            // get position Info
            positionInfo = positionInfos[positionIndexes[i]];
            _safeClosePosition(ledgerType, epoch, positionIndexes[i], positionInfo, roundInfo, userGlobalInfo);
        }
    }

    /**
     * @notice Claim a batch of incentive claimable positions
     * @param ledgerType: Ledger type
     * @param epoch: Epoch of the ledger
     * @param positionIndexes: Position indexes of the user
     */

    //批量取出激励金额
    function batchClaimPositionIncentiveReward(
        uint256 ledgerType,
        uint256 epoch,
        uint256[] calldata positionIndexes
    ) external notContract nonReentrant {
        require(ledgerType < 6, "Invalid ledger type");
        require(epoch < currentEpochs[ledgerType], "Epoch not finished");

        // get position infos
        PositionInfo[] storage positionInfos = roundLedgers[ledgerType][epoch][msg.sender];

        // get roundInfo
        RoundInfo storage roundInfo = roundInfos[ledgerType][epoch];

        // get user round info
        uint256 userRoundIndex = roundInfo.ledgerRoundToUserRoundIndex[msg.sender];
        UserRoundInfo storage userRoundInfo = userRoundsInfos[ledgerType][msg.sender][userRoundIndex];

        // position info placeholder
        PositionInfo storage positionInfo;

        // collect payout
        uint256 payoutAmount;
        uint256 positionIndex;
        uint256 positionIndexesLength = positionIndexes.length;
        uint256 positionInfosLength = positionInfos.length;
        for (uint256 i = 0; i < positionIndexesLength; ++i) {
            positionIndex = positionIndexes[i];
            require(positionIndex < positionInfosLength, "Invalid position index");
            // get position Info
            positionInfo = positionInfos[positionIndex];
            require(positionInfo.incentiveClaimable, "Position not eligible");
            // update positionInfo
            payoutAmount += _safeProcessIncentiveAmount(positionInfo, roundInfo);
        }

        // transfer
        {
            (bool success, ) = msg.sender.call{value: payoutAmount}("");
            require(success, "Transfer failed.");
        }

        // update userRoundInfo
        userRoundInfo.totalIncentiveClaimedAmount += payoutAmount;
        emit IncentiveClaimed(msg.sender, payoutAmount);
    }

    /**
     * @notice Report a batch users' sales
     * @param users: list of users
     */
    function batchReportSales(address[] calldata users) external {
        uint256 usersLength = users.length;
        for (uint256 i = 0; i < usersLength; ++i) {
            _safeReportSales(users[i]);
        }
    }

    /**
     * @notice Claim referrer reward
     * @param referrer: referrer address
     */
    //转出代理资金
    function claimReferrerReward(address referrer) external notContract nonReentrant {
        require(referrer != address(0), "Invalid referrer address");

        // get user global info
        UserGlobalInfo storage userGlobalInfo = userGlobalInfos[referrer];

        // get claimable amount
        uint256 claimableAmount = userGlobalInfo.totalReferrerReward - userGlobalInfo.referrerRewardClaimed;

        require(claimableAmount > 0, "No claimable amount");

        // update state
        userGlobalInfo.referrerRewardClaimed += claimableAmount;

        // do transfer
        {
            (bool success, ) = referrer.call{value: claimableAmount}("");
            require(success, "Transfer failed.");
        }

        // emit event
        emit ReferrerRewardClaimed(referrer, claimableAmount);
    }

    function getLinkedPositionInfo(
        uint256 ledgerType,
        uint256 epoch,
        uint256 cursor,
        uint256 size
    ) external view returns (LinkedPosition[] memory, uint256) {
        uint256 length = size;
        uint256 positionCount = roundInfos[ledgerType][epoch].totalPositionCount;
        if (cursor + length > positionCount) {
            length = positionCount - cursor;
        }
        LinkedPosition[] memory linkedPositions = new LinkedPosition[](length);
        RoundInfo storage roundInfo = roundInfos[ledgerType][epoch];
        for (uint256 i = 0; i < length; ++i) {
            linkedPositions[i] = roundInfo.linkedPositions[cursor + i];
        }
        return (linkedPositions, cursor + length);
    }

    function getUserRounds(
        uint256 ledgerType,
        address user,
        uint256 cursor,
        uint256 size
    ) external view returns (UserRoundInfo[] memory, uint256) {
        uint256 length = size;
        uint256 roundCount = userRoundsInfos[ledgerType][user].length;
        if (cursor + length > roundCount) {
            length = roundCount - cursor;
        }

        UserRoundInfo[] memory userRoundInfos = new UserRoundInfo[](length);
        for (uint256 i = 0; i < length; ++i) {
            userRoundInfos[i] = userRoundsInfos[ledgerType][user][cursor + i];
        }

        return (userRoundInfos, cursor + length);
    }

    function getUserRoundsLength(uint256 ledgerType, address user) external view returns (uint256) {
        return userRoundsInfos[ledgerType][user].length;
    }

    function getUserRoundLedgers(
        uint256 ledgerType,
        uint256 epoch,
        address user,
        uint256 cursor,
        uint256 size
    ) external view returns (PositionInfo[] memory, uint256) {
        uint256 length = size;
        uint256 positionCount = roundLedgers[ledgerType][epoch][user].length;
        if (cursor + length > positionCount) {
            length = positionCount - cursor;
        }

        PositionInfo[] memory positionInfos = new PositionInfo[](length);
        for (uint256 i = 0; i < length; ++i) {
            positionInfos[i] = roundLedgers[ledgerType][epoch][user][cursor + i];
        }

        return (positionInfos, cursor + length);
    }

    function getUserRoundLedgersLength(
        uint256 ledgerType,
        uint256 epoch,
        address user
    ) external view returns (uint256) {
        return roundLedgers[ledgerType][epoch][user].length;
    }

    function getChildren(
        address user,
        uint256 cursor,
        uint256 size
    ) external view returns (address[] memory, uint256) {
        uint256 length = size;
        uint256 childrenCount = children[user].length;
        if (cursor + length > childrenCount) {
            length = childrenCount - cursor;
        }

        address[] memory _children = new address[](length);
        for (uint256 i = 0; i < length; ++i) {
            _children[i] = children[user][cursor + i];
        }

        return (_children, cursor + length);
    }

    function getLedgerRoundToUserRoundIndex(
        uint256 ledgerType,
        uint256 epoch,
        address user
    ) external view returns (uint256) {
        return roundInfos[ledgerType][epoch].ledgerRoundToUserRoundIndex[user];
    }

    function getChildrenLength(address user) external view returns (uint256) {
        return children[user].length;
    }

    function getUserDepartSalesAndLevel(address user) external view returns (uint256, uint8) {
        UserGlobalInfo storage userGlobalInfo = userGlobalInfos[user];
        return (userGlobalInfo.sales - userGlobalInfo.maxChildrenSales, userGlobalInfo.salesLevel);
    }

    /**
     * @notice close a given position
     * @param ledgerType: ledger type
     * @param epoch: epoch of the ledger
     * @param positionIndex: position index of the user
     * @param positionInfo: storage of the position info
     * @param roundInfo: storage of the round info
     */
    function _safeClosePosition(
        uint256 ledgerType,
        uint256 epoch,
        uint256 positionIndex,
        PositionInfo storage positionInfo,
        RoundInfo storage roundInfo,
        UserGlobalInfo storage userGlobalInfo
    ) internal {
        //*** positionInfo 初始化的时候 withdrawnAmount 为zero */
        require(positionInfo.withdrawnAmount == 0, "Position already claimed");
        //盒子过期或者此纪元结束
        require(positionInfo.expiryTime <= block.timestamp || roundInfo.stopLoss, "Position not expired");

        // get user round info from storage
        uint256 targetRoundInfoIndex = roundInfo.ledgerRoundToUserRoundIndex[msg.sender];
        UserRoundInfo storage userRoundInfo = userRoundsInfos[ledgerType][msg.sender][targetRoundInfoIndex];

        // calculate the amount to withdraw
        uint256 payoutAmount; //退回金额
        uint256 principalAmount = (positionInfo.amount * PRINCIPAL_RATIO) / PRICE_PRECISION;

        // get back the principal amount
        payoutAmount += principalAmount;

        // update roundInfo
        //更新轮数信息
        roundInfo.currentPositionCount -= 1;
        roundInfo.currentPrincipalAmount -= principalAmount;
        //此纪元正在进行
        if (!roundInfo.stopLoss) {
            // calculate expected invest return amount
            // how many days passed
            //计算预期收益 静态投资
            uint256 daysPassed;
            if (ledgerType == 0) {
                // 1 day
                daysPassed = (block.timestamp - positionInfo.openTime);
            } else {
                daysPassed = (positionInfo.expiryTime - positionInfo.openTime);
            }
            //额外获得增益
            uint256 expectedInvestReturnAmount = (positionInfo.amount * positionInfo.investReturnRate * daysPassed) /
                PRICE_PRECISION /
                TIME_UNIT;

            // calculate the amount should be paid back from invest pool
            // 35% to total amount + expected return amount
            //除开基础的百分之65加上额外收益应该获得的总数
            uint256 investReturnAmount = positionInfo.amount - principalAmount + expectedInvestReturnAmount;

            // compare if current invest pool has enough amount
            //如果本轮资金不够，则尽量退还
            if (roundInfo.currentInvestAmount < investReturnAmount) {
                // not enough, then just pay back the current invest pool amount
                investReturnAmount = roundInfo.currentInvestAmount;
                roundInfo.currentInvestAmount = 0;
            } else {
                // update round info
                unchecked {
                    roundInfo.currentInvestAmount -= investReturnAmount;
                }
            }

            // check round is stop loss
            //如果当前纪元对应等级的静态投资投资金额为0的话,默认爆仓,暂停本轮,重开一轮
            if (roundInfo.currentInvestAmount == 0) {
                roundInfo.stopLoss = true;
                currentEpochs[ledgerType] += 1;
                _refillStock(ledgerType);
                emit NewRound(currentEpochs[ledgerType], ledgerType);
            }

            // update payout amount
            //更新能得到的金额
            payoutAmount += investReturnAmount;
            //此轮静态投资能够获取的投资收益(初始为35%)
            // update positionInfo
            positionInfo.investReturnAmount = investReturnAmount;
        }
       //激励奖励
        uint256 incentiveAmount = 0;
        // calculate incentive amount if eligible
        //只有当本轮暂停后,另外盒子处于最后六十个能够得到另外收益
        if (roundInfo.stopLoss && positionInfo.incentiveClaimable) {
            incentiveAmount = _safeProcessIncentiveAmount(positionInfo, roundInfo);

            // update payout amount
            payoutAmount += incentiveAmount;

            // update incentive info to storage
            userRoundInfo.totalIncentiveClaimedAmount += incentiveAmount;

            emit IncentiveClaimed(msg.sender, incentiveAmount);
        }

        // update user round info
        //用户此纪元对应等级的总共回撤金额
        userRoundInfo.totalWithdrawnAmount += payoutAmount;
        //减去基本保障
        userRoundInfo.currentPrincipalAmount -= principalAmount;

        // update positionInfo
        //修改此静态投资的回撤金额
        positionInfo.withdrawnAmount = payoutAmount;

        // accumulate user's boost credit
        //如果用户在此静态投资的期间亏了钱,那么用户能够获得boostCredit增加收益的凭证，保障下一次购买静态投资的时候在一定程度上能够多获得0.5%的利率
        if (payoutAmount - incentiveAmount < positionInfo.amount) {
            userGlobalInfo.boostCredit += positionInfo.amount;
        }

        // do transfer
        {
            //给钱
            (bool success, ) = msg.sender.call{value: payoutAmount}("");
            require(success, "Transfer failed.");
        }

        // emit event
        emit PositionClosed(msg.sender, ledgerType, epoch, positionIndex, payoutAmount);
    }

    /**
     * @notice process current round's fund target and return the updated invest return rate
     * @param roundInfo: storage of the round info
     */
    //获取投资回报比例
    function _safeProcessFundTargetGetInvestReturnRate(RoundInfo storage roundInfo, uint256 ledgerType)
        internal
        returns (uint256)
    {
        FundTarget storage fundTarget = roundInfo.fundTarget;
        uint256 targetAmount = fundTarget.amount;
        uint256 achievedAmount = fundTarget.achievedAmount;
        // this is amount of total locked position
        uint256 currentTotalAmount_d6 = roundInfo.currentPrincipalAmount * PRICE_PRECISION;

        // process target fund
        {
            // check if this is the first time to process fund target
            if (fundTarget.lastCheckTime == 0) {
                // first check will use default parameter
                targetAmount = DEFAULT_TARGET_AMOUNTS[ledgerType];

                // update check time and target amount to storage
                fundTarget.lastCheckTime = block.timestamp;
                fundTarget.amount = targetAmount;
            } else {
                // check if over 24 hours since last check
                if (block.timestamp - fundTarget.lastCheckTime > TIME_UNIT) {
                    // recalculate target amount
                    targetAmount =
                        (((currentTotalAmount_d6 * 361) / 1000 / PRINCIPAL_RATIO - roundInfo.currentInvestAmount) *
                            PRICE_PRECISION) /
                        260000;

                    // update check time and target amount to storage
                    fundTarget.lastCheckTime = block.timestamp;
                    fundTarget.amount = targetAmount;
                    // reset achieved amount
                    fundTarget.achievedAmount = 0;
                    // reset achievedAmount in memory as well, because this will be the first position after adjusting the FundTarget
                    achievedAmount = 0;
                }
            }
            // update achieved amount in storage
            fundTarget.achievedAmount += msg.value;
        }

        // calculate return rate
        // notice: no need to include current invest amount
        if (achievedAmount <= targetAmount) {
            return DEFAULT_INVEST_RETURN_RATE;
        }

        // decrease 0.05% per 20% over target amount till 0.3%
        uint256 ratioDiff = (achievedAmount * PRICE_PRECISION) / targetAmount - PRICE_PRECISION;
        uint256 times = ratioDiff / (200000) + 1;
        if (ratioDiff % (200000) == 0) {
            times -= 1;
        }
        if (times > 14) {
            times = 14;
        }

        return DEFAULT_INVEST_RETURN_RATE - (times * 500);
    }

    /**
     * @notice process positionInfo and return incentive amount
     * @param positionInfo: storage of the position info
     * @param roundInfo: storage of the round info
     */
    //计算另外得到的激励金额
    function _safeProcessIncentiveAmount(PositionInfo storage positionInfo, RoundInfo storage roundInfo)
        internal
        returns (uint256)
    {
        // calculate incentive amount
        //***将所有的激励奖励根据最后六十位的盒子金额总数分配所占有的比重 */
        uint256 incentiveAmount = (positionInfo.amount * roundInfo.totalPositionAmount * INCENTIVE_RATIO) /
            roundInfo.incentiveSnapshot /
            PRICE_PRECISION;

        // with PRICE_PRECISION is due to the precision of division may result in a few wei left over
        //当前用于真正激励的金额如果小于理论应该得到的，那么就尽量满足
        if (roundInfo.currentIncentiveAmount < incentiveAmount + PRICE_PRECISION) {
            // clean up incentive amount
            incentiveAmount = roundInfo.currentIncentiveAmount;
            roundInfo.currentIncentiveAmount = 0;
        } else {
            roundInfo.currentIncentiveAmount -= incentiveAmount;
        }

        // this position is no longer eligible for incentive
        positionInfo.incentiveClaimable = false;

        // update positionInfo
        //激励奖励,奖励最后六十个
        positionInfo.incentiveAmount = incentiveAmount;

        return incentiveAmount;
    }

    /**
     * @notice process user's level info and return the current level
     * @param currentLevel: user current level
     * @param user: user address
     * @param currentSales: user current sales
     * @param userGlobalInfo: storage of the user global info
     */
    //所处社区的等级
    function _safeProcessSalesLevel(
        uint8 currentLevel,
        address user,
        uint256 currentSales,
        UserGlobalInfo storage userGlobalInfo
    ) internal returns (uint8) {
        uint8 newLevel = _getSalesToLevel(currentSales);
        if (newLevel > currentLevel) {
            userGlobalInfo.salesLevel = newLevel;
            emit SalesLevelUpdated(user, newLevel);
        } else {
            newLevel = currentLevel;
        }
        return newLevel;
    }

    /**
     * @notice report user's sales and update its referrer sales level
     * @param user: user address
     */
    function _safeReportSales(address user) internal {
        UserGlobalInfo storage userGlobalInfo = userGlobalInfos[user];
        address referrer = userGlobalInfo.referrer;
        uint256 userSales = userGlobalInfo.sales;
        uint256 userReportedSales = userGlobalInfo.reportedSales;

        // get user's un-reported sales
        uint256 unreportedSales = userSales - userReportedSales;

        if (unreportedSales > 0) {
            // get referrer global info from storage
            UserGlobalInfo storage referrerGlobalInfo = userGlobalInfos[referrer];
            // fill up the sales to the referrer
            referrerGlobalInfo.sales += unreportedSales;
            // update user's reported sales
            userGlobalInfo.reportedSales = userSales;

            // all reported sales + user's own contributed position will be current user's final sales
            userSales += userGlobalInfo.totalPositionAmount;
            // current referrer's max children sales
            uint256 maxChildrenSales = referrerGlobalInfo.maxChildrenSales;
            // update max children sales if needed
            if (userSales > maxChildrenSales) {
                // referrer's max children sales is updated
                referrerGlobalInfo.maxChildrenSales = userSales;
                // update cache of max children sales
                maxChildrenSales = userSales;
            }
            // process referrer's sales level
            _safeProcessSalesLevel(
                referrerGlobalInfo.salesLevel,
                referrer,
                referrerGlobalInfo.sales - maxChildrenSales, // sales for level calculation is sales - max children sales
                referrerGlobalInfo
            );
        }
    }

    /**
     * @notice distribute referrer reward
     * @param user: user address (被推荐人的地址)
     * @param referrerAmount: total amount of referrer reward (推荐人应当动态投资所获得的回报)
     */
    //分配推荐人的奖励
    function _distributeReferrerReward(address user, uint256 referrerAmount) internal virtual {
        UserGlobalInfo storage userGlobalInfo = userGlobalInfos[user];
        UserGlobalInfo storage referrerGlobalInfo;
        //被推荐人静态投资的数额(数额和静态投资类型无关)
        uint256 positionAmount = msg.value;

        // init all local variables as a search struct
        ReferrerSearch memory search;
        // 怎样去理解 basesaleslevel
        search.baseSalesLevel = 0;
        //当前梯度
        search.currentReferrer = userGlobalInfo.referrer;
        //百分之八十的部落奖
        search.levelDiffAmount = (referrerAmount * 80) / 100;
        search.leftLevelDiffAmount = search.levelDiffAmount;
        //不同梯度的分配金额差异
        search.levelDiffAmountPerLevel = search.levelDiffAmount / 12;
        //百分之二十的精灵奖
        search.levelSearchAmount = referrerAmount - search.levelDiffAmount;
        search.leftLevelSearchAmount = search.levelSearchAmount;
        //精灵奖
        search.levelSearchAmountPerReferrer = search.levelSearchAmount / 10;
        //更新用户的静态投资金额(这里修改了用户全局信息)
        search.currentUserTotalPosAmount = userGlobalInfo.totalPositionAmount + positionAmount;
        userGlobalInfo.totalPositionAmount = search.currentUserTotalPosAmount;
        search.currentUser = user;
       
        while (search.depth < MAX_SEARCH_DEPTH) {
            // stop if current referrer is the root
            if (search.currentReferrer == address(0)) {
                break;
            }

            // this position does not counted as reported sales for first user himself
            if (search.depth > 0) userGlobalInfo.reportedSales += positionAmount;

            // cache current user information
            search.currentUserSales = userGlobalInfo.sales;
            search.currentUserReportedSales = userGlobalInfo.reportedSales;

            // cache current referrer information
            referrerGlobalInfo = userGlobalInfos[search.currentReferrer];

            // update referrer sales
            {
                search.currentReferrerSales = referrerGlobalInfo.sales;
                // add current sales to current referrer
                search.currentReferrerSales += positionAmount; //****  KEY */
                // check unreported sales
                //如果当前user的reportedSales小于sales
                if (search.currentUserReportedSales < search.currentUserSales) {
                    // update referrerSales to include unreported sales
                    //referrersales加上当前user的unreported
                    search.currentReferrerSales += search.currentUserSales - search.currentUserReportedSales;
                    // update current node storage for reported sales
                    userGlobalInfo.reportedSales = search.currentUserSales;
                }
                // update sales for current referrer
                referrerGlobalInfo.sales = search.currentReferrerSales;
            }

            // update referrer max children sales //***** KEY */
            //是否更新userglobal的maxchildrenSales
            {
                // add current user's total position amount to current user's sales
                search.currentUserSales += search.currentUserTotalPosAmount;
                // check referrer's max child sales
                search.currentReferrerMaxChildSales = referrerGlobalInfo.maxChildrenSales;
                if (search.currentReferrerMaxChildSales < search.currentUserSales) {
                    // update max child sales
                    referrerGlobalInfo.maxChildrenSales = search.currentUserSales;
                    search.currentReferrerMaxChildSales = search.currentUserSales;
                }
            }

            // process referrer's sales level
            // @notice: current referrer sales level should ignore its max child sales
            search.currentReferrerLevel = _safeProcessSalesLevel(
                referrerGlobalInfo.salesLevel,
                search.currentReferrer,
                search.currentReferrerSales - search.currentReferrerMaxChildSales,
                referrerGlobalInfo
            );

            // start level diff calculation
            //部落奖分配规则
            if (!search.levelDiffDone) {
                // compare the current referrer's level with the base sales level
                if (search.currentReferrerLevel > search.baseSalesLevel) {
                    // level diff
                    search.currentLevelDiff = search.currentReferrerLevel - search.baseSalesLevel;

                    // update base level
                    search.baseSalesLevel = search.currentReferrerLevel;

                    // calculate the referrer amount
                    //与上一级的社区等级差异计算currentReferrerAmount(所相差的社区奖励)
                    search.currentReferrerAmount = search.currentLevelDiff * search.levelDiffAmountPerLevel;

                    // check left referrer amount
                    if (search.currentReferrerAmount + PRICE_PRECISION > search.leftLevelDiffAmount) {
                        search.currentReferrerAmount = search.leftLevelDiffAmount;
                    }

                    // update referrer's referrer amount
                    referrerGlobalInfo.totalReferrerReward += search.currentReferrerAmount;
                    emit ReferrerRewardAdded(search.currentReferrer, search.currentReferrerAmount, 0);

                    unchecked {
                        search.leftLevelDiffAmount -= search.currentReferrerAmount;
                    }

                    if (search.leftLevelDiffAmount == 0) {
                        search.levelDiffDone = true;
                    }
                }
            } 
            //精灵奖分配规则
            if (!search.levelSearchDone) {
                // level search use referrer's real level
                search.levelSearchStep = _getLevelToLevelSearchStep(
                    _getSalesToLevel(search.currentReferrerSales - search.currentReferrerMaxChildSales)
                );

                if (search.numLevelSearchCandidate + 1 <= search.levelSearchStep) {
                    search.numLevelSearchCandidate += 1;

                    // check left referrer amount
                    if (search.levelSearchAmountPerReferrer + PRICE_PRECISION > search.leftLevelSearchAmount) {
                        search.levelSearchAmountPerReferrer = search.leftLevelSearchAmount;
                    }

                    // update referrer's referrer amount
                    referrerGlobalInfo.totalReferrerReward += search.levelSearchAmountPerReferrer;
                    emit ReferrerRewardAdded(search.currentReferrer, search.levelSearchAmountPerReferrer, 1);

                    unchecked {
                        search.leftLevelSearchAmount -= search.levelSearchAmountPerReferrer;
                    }

                    if (search.leftLevelSearchAmount == 0) {
                        search.levelSearchDone = true;
                    }
                }
            }

            search.currentUser = search.currentReferrer;
            search.currentReferrer = referrerGlobalInfo.referrer;

            userGlobalInfo = referrerGlobalInfo;
            search.currentUserTotalPosAmount = userGlobalInfo.totalPositionAmount;

            unchecked {
                search.depth += 1;
            }
        }

        // check residual referrer amount
        //没用完就给用户
        if (search.leftLevelDiffAmount > 0) {
            userGlobalInfos[user].totalReferrerReward += search.leftLevelDiffAmount;
            emit ReferrerRewardAdded(user, search.leftLevelDiffAmount, 0);
        }
        if (search.leftLevelSearchAmount > 0) {
            userGlobalInfos[user].totalReferrerReward += search.leftLevelSearchAmount;
            emit ReferrerRewardAdded(user, search.leftLevelSearchAmount, 1);
        }
    }

    /**
     * @notice get sales level from sales amount
     * @param amount: sales amount
     */
    //根据销售额算等级
    function _getSalesToLevel(uint256 amount) internal pure virtual returns (uint8) {
        /* istanbul ignore else  */
        if (amount < 10000 ether) {
            return 0;
        } else if (amount < 100000 ether) {
            return 1;
        } else if (amount < 400000 ether) {
            return 2;
        } else if (amount < 800000 ether) {
            return 3;
        } else if (amount < 1500000 ether) {
            return 4;
        } else if (amount < 3000000 ether) {
            return 5;
        } else if (amount < 6000000 ether) {
            return 6;
        } else if (amount < 12000000 ether) {
            return 7;
        } else if (amount < 30000000 ether) {
            return 8;
        } else if (amount < 60000000 ether) {
            return 9;
        } else if (amount < 120000000 ether) {
            return 10;
        } else if (amount < 240000000 ether) {
            return 11;
        }
        return 12;
    }

    /**
     * @notice level search step from level
     * @param level: sales level (0-12)
     */
    //精灵等级
    function _getLevelToLevelSearchStep(uint8 level) internal pure returns (uint8) {
        unchecked {
            if (level < 5) return level * 2;
        }
        return 10;
    }
}