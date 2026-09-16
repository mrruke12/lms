create type user_type as enum ('student', 'teacher');
create type assessment_type as enum('none', 'manual', 'predefined', 'tests');

create type lesson_status as enum('draft', 'editing', 'published', 'archived');
create type attempt_status as enum('active', 'submitted', 'completed', 'overdue', 'archived');
create type submission_status as enum('pending', 'evaluated');
create type assessment_status as enum('pending', 'partial', 'evaluated');
create type enrollment_status as enum('active', 'archived');

create type degree_type as enum('associate', 'bachelor', 'master', 'doctor');

create table users (
    id UUID primary key default uuidv7(),
    type user_type not null,
    login varchar(32) not null unique,
    password_hash varchar(256) not null,
    name varchar(32) not null,
    surname varchar(32) not null,
    patronymic varchar(32) not null
);

create table lessons (
    id UUID primary key default uuidv7(),
    author_id UUID references users(id) on delete set null, -- set status to archive in delete user transaction
    version UUID not null,
    name varchar(256) not null,
    status lesson_status not null,

    constraint uq_lesson_version unique (id, version)
);

create table element_types (
    id serial primary key,
    name varchar(256) not null unique
);

create table elements (
    id UUID primary key default uuidv7(),
    lesson_id UUID not null references lessons(id) on delete cascade, 
    parent_id UUID references elements(id) on delete cascade,
    type_id integer references element_types(id) on delete set null,
    assessment assessment_type not null,
    config jsonb not null
);

create table revisions (
    id UUID primary key default uuidv7(),
    element_id UUID not null,
    lesson_id UUID not null references lessons(id) on delete cascade,
    parent_id UUID, 
    type_id integer references element_types(id) on delete set null,
    assessment assessment_type not null,
    config jsonb not null
);

create table attempts (
    id UUID primary key default uuidv7(),
    lesson_id UUID not null references lessons(id) on delete cascade,
    lesson_version UUID not null, 
    student_id UUID not null references users(id) on delete cascade,
    status attempt_status not null
);

create table submission (
    id UUID primary key default uuidv7(),
    attempt_id UUID not null references attempts(id) on delete cascade,
    element_id UUID references elements(id),
    cap integer not null,
    score integer not null,
    status submission_status not null
);

create table assessment (
    id UUID primary key default uuidv7(),
    attempt_id UUID not null references attempts(id) on delete cascade,
    status assessment_status not null,
    cap integer not null,
    grade integer not null
);

create table groups (
    id UUID primary key default uuidv7(),
    name varchar(256) not null,
    degree degree_type not null,
    entry_year integer not null,
    graduation_year integer not null
);

create table group_enrollments (
    id UUID primary key default uuidv7(),
    user_id UUID not null references users(id) on delete cascade,
    group_id UUID not null references groups(id) on delete cascade, 
    status enrollment_status not null
);


create table lesson_enrollments (
    id UUID primary key default uuidv7(),
    user_id UUID not null references users(id) on delete cascade,
    group_id UUID not null references lessons(id) on delete cascade, 
    status enrollment_status not null
);