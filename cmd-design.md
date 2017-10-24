# 命令设计

## 列出帮助信息

    agenda help // 列出所有帮助信息
    agenda help [command] // 列出指定命令的帮助信息

## 用户注册

    agenda register --un=UserName --pw=password --email=a@xxx.com --phone=xxxxxx

## 用户登录

    agenda login --un=UserName --pw=password

## 用户登出

    agenda logout

## 用户查询

    agenda query --un=UserName

## 用户删除

    agenda delUser

## 创建会议

    agenda cm --title=meeting --pr=participator --st=start_time --et=end_time

## 增删会议参与者

    agenda addPr --title=meeting --pr=participator // 增加会议参与者
    agenda delPr --title=meeting --pr=participator // 删除会议参与者

## 查询会议

    agenda qm --st=start_time  --et=end_time

## 取消会议

    agenda delm --title=meeting

## 退出会议

    agenda exitm --title=meeting

## 清空会议

    agenda clear
