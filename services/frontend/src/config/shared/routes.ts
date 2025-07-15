import {
  IconBookFilled,
  IconBook,
  IconHome,
  IconHomeFilled,
  IconUser,
  IconMessage,
  IconMessageFilled,
  IconVideo,
  IconVideoFilled,
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
      path: '/dashboard/videos',
      name: 'Videos',
    },
    {
      path: '/dashboard/tests',
      name: 'Tests',
    },
    {
      path: '/dashboard/codelab',
      name: 'Code Lab',
    },
  ],
  teacherRoutes: [
    {
      path: '/dashboard',
      name: 'Student View ',
      icon: IconUser,
      iconSelect: IconHomeFilled,
    },
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
      name: 'Videos',
      path: '/teacher/videos',
      icon: IconVideo,
      iconSelect: IconVideoFilled,
    },
    {
      path: '/teacher/tests',
      name: 'Tests',
      icon: IconMessage,
      iconSelect: IconMessageFilled,
    },
    {
      path: '/teacher/codelab',
      name: 'Code Lab',
      icon: IconVideo,
      iconSelect: IconVideoFilled,
    },
  ],
};
