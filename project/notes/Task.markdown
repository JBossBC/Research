# Task

## Question

class: ①input-group input group-sm ②input-group-prepend ③form-control ④btn btn-primary btn-sm row-icon

## 需要的另外的功能:

+ 查询

    根据用户查找服务器

     根据ip地址查找服务器

     根据服务器ID查找服务器

     根据所在组织查找服务器

前端数据:

参数植入 :   
    
       <div class="wid_15">
       <div class="input-group input group-sm">
    		<div class="input-group-prepend">
      <span class="input-group-text">所属用户</span>
    		</div>
    		<input type="text" class="form-control">
    	   </div>
    	</div>

事件驱动:

    <button type="button" class="btn btn-primary btn-sm row-icon">
    			<div class="svg-icon">
    				<svg viewBox="0 0 1024 1024" width="100%" height="100%">
    					<path d="M756.7 183.4c-158.3-158.3-415-158.3-573.3 0s-158.3 415 0 573.3c142 142 362.9 156.2 521.2 43.5L844 939.6c26.3 26.3 69.3 26.3 95.6 0s26.3-69.3 0-95.6L800.2 704.7c112.7-158.4 98.5-379.3-43.5-521.3z m-95.5 477.8c-105.4 105.4-276.8 105.4-382.2 0S173.6 384.3 279 279s276.8-105.4 382.2 0 105.4 276.8 0 382.2z"></path>
    				</svg>
    			</div>
    			查询
    		</button>



VM table:




  Name|Code|DataType|Length|Precision|P|F|M
    -|-|-|-|-|-|-|-
  所属用户|USER_ID|VARCHAR(64)|64|			


所属用户	USER_ID	VARCHAR(64)	64				




