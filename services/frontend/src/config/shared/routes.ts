import {
  IconBookFilled,
  IconBook,
  IconHome,
  IconHomeFilled,
  IconMessageReport,
  IconUser,
  IconUserFilled,
  IconMessageReportFilled,
  IconMessage,
  IconMessageFilled,
} from '@tabler/icons-react';

export const routesConfig = {
  studentRoutes: [
    {
      path: '/dashboard',
      name: 'Dashboard',
    },
    {
      path: '/dashboard/courses',
      name: 'Courses',
    },
    {
      path: '/dashboard/post',
      name: 'Posts',
    },
  ],
  teacherRoutes: [
    {
      path: '/teacher',
      name: 'Dashboard',
      icon: IconHome,
      iconSelect: IconHomeFilled,
    },
    {
      path: '/teacher/courses',
      name: 'Courses',
      icon: IconBook,
      iconSelect: IconBookFilled,
    },
    {
      name: 'Students',
      path: '/teacher/students',
      icon: IconUser,
      iconSelect: IconUserFilled,
    },
    {
      path: '/teacher/post',
      name: 'Posts',
      icon: IconMessage,
      iconSelect: IconMessageFilled,
    },
    {
      icon: IconMessageReport,
      iconSelect: IconMessageReportFilled,
      name: 'Reports',
      path: '/teacher/reports',
    },
  ],
};
