# mybatis-gorm
这是一个基于gorm的封装，对于gorm的原生sql语法做了一些优化，支持形如select * from users where email like #{email} and name like "王%" order by created_at limit #{pageSize}的sql语法
