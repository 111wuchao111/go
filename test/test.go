 $.ajax({

               url ：'/user/get-online-users',  //后台处理程序

               type:'get',    //数据发送方式

               dataType:'json',  //接受数据格式

               //data:params,  //要传递的数据

               success:function(data) {  
						alert(data)  
				   },  
				error : function() {   
				     alert("异常！");  
				   }  

               });

        });