CREATE USER 'sibcoder'@'%' IDENTIFIED BY '210104Sib!';

create Database project;
use project;

SET NAMES 'utf8';

create table projects
(
    id          int  not null AUTO_INCREMENT
        primary key,
    name        varchar(100)  null,
    start_date  date          null,
    end_date    date          null,
    colour      varchar(100)  null,
    description varchar(1000) null
) DEFAULT CHARSET=utf8;


create table employees
(
    id         int not null AUTO_INCREMENT
        primary key,
    login      varchar(100) not null,
    password   varchar(100) not null,
    first_name varchar(100) null,
    last_name  varchar(100) null,
    is_admin   tinyint(1)   null
) DEFAULT CHARSET=utf8;


create table tasks
(
    id          int not null AUTO_INCREMENT
        primary key,
    name        varchar(100)  null,
    start_date  date          null,
    end_date    date          null,
    project_id  int           null,
    description varchar(1000) null,
    is_finished boolean not null default 0,

    constraint tasks_projects_fk
        foreign key (project_id) references projects (id)
) DEFAULT CHARSET=utf8;


create table task_for_emp
(
    task_id int not null,
    emp_id  int not null,
    primary key (task_id, emp_id),
    constraint task_for_emp_ibfk_1
        foreign key (task_id) references tasks (id),
    constraint task_for_emp_ibfk_2
        foreign key (emp_id) references employees (id)
) DEFAULT CHARSET=utf8;

create index emp_id
    on task_for_emp (emp_id);



alter table projects add text_colour varchar(100) null;


INSERT INTO  employees (login, password, first_name, last_name, is_admin) VALUES ('sib_coder', '210104', 'Sib', 'Coder', 1);
INSERT INTO  employees (id ,login, password, first_name, last_name, is_admin) VALUES (2,'RomanL', '1234', 'Roman', 'Lider', 0);
INSERT INTO  employees (login, password, first_name, last_name, is_admin) VALUES ('crow', '1234', 'Ksenia', 'Andrianova', 0);
INSERT INTO  employees (login, password, first_name, last_name, is_admin) VALUES ('Andr', '1234', 'Andrei', 'Davidkiv', 0);
INSERT INTO  employees (login, password, first_name, last_name, is_admin) VALUES ('sibears', '04051970', 'Si', 'Bears', 0);


INSERT INTO  projects (id,name, start_date, end_date, colour, description, text_colour) VALUES (1 ,'Менеджер Проектов', STR_TO_DATE('2022-11-28', '%Y-%m-%d'), STR_TO_DATE('2022-12-28', '%Y-%m-%d'), '#f9f06b', 'Создание менеджера проектов', '#000000');

INSERT INTO  tasks (id ,name, start_date, end_date, project_id, description, is_finished) VALUES (1,'Сдать 1 фазу', STR_TO_DATE('2022-11-28', '%Y-%m-%d'), STR_TO_DATE('2022-12-28', '%Y-%m-%d'), 1, 'Создание менеджера проектов. Задачи 1 фазы.', 0);
INSERT INTO  tasks (name, start_date, end_date, project_id, description, is_finished) VALUES ('Сдать 2 фазу', STR_TO_DATE('2022-11-28', '%Y-%m-%d'), STR_TO_DATE('2022-12-28', '%Y-%m-%d'), 1, 'Создание менеджера проектов. Задачи 2 фазы.', 0);
INSERT INTO  tasks (name, start_date, end_date, project_id, description, is_finished) VALUES ('Сдать 3 фазу', STR_TO_DATE('2022-11-28', '%Y-%m-%d'), STR_TO_DATE('2022-12-28', '%Y-%m-%d'), 1, 'Создание менеджера проектов. Задачи 3 фазы.', 0);

INSERT INTO task_for_emp (task_id, emp_id) VALUES(1,2);


GRANT ALL PRIVILEGES ON *.* TO 'sibcoder'@'%' WITH GRANT OPTION;

FLUSH PRIVILEGES;