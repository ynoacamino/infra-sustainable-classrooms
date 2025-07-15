'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/ui/card';
import { Badge } from '@/ui/badge';
import { Avatar, AvatarFallback, AvatarImage } from '@/ui/avatar';
import { Button } from '@/ui/button';
import {
  Edit,
  Mail,
  Phone,
  User,
  GraduationCap,
  Briefcase,
} from 'lucide-react';
import type { CompleteProfile } from '@/types/profiles/models';
import Link from 'next/link';
import { formatDate } from '@/lib/shared/utils';

interface ProfileViewProps {
  profile: CompleteProfile;
}

export function ProfileView({ profile }: ProfileViewProps) {
  const getInitials = (firstName: string, lastName: string) => {
    return `${firstName.charAt(0)}${lastName.charAt(0)}`.toUpperCase();
  };

  const getRoleColor = (role: string) => {
    return role === 'teacher'
      ? 'bg-blue-100 text-blue-800'
      : 'bg-green-100 text-green-800';
  };

  const getRoleIcon = (role: string) => {
    return role === 'teacher' ? (
      <Briefcase className="w-4 h-4" />
    ) : (
      <GraduationCap className="w-4 h-4" />
    );
  };

  return (
    <div className="w-full max-w-4xl mx-auto p-4 space-y-6">
      {/* Header with basic info */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <div className="flex items-center space-x-4">
            <Avatar className="w-20 h-20">
              <AvatarImage
                src={profile.avatar_url}
                alt={`${profile.first_name} ${profile.last_name}`}
              />
              <AvatarFallback className="text-lg">
                {getInitials(profile.first_name, profile.last_name)}
              </AvatarFallback>
            </Avatar>
            <div className="space-y-1">
              <h1 className="text-2xl font-bold">
                {profile.first_name} {profile.last_name}
              </h1>
              <div className="flex items-center space-x-2">
                {getRoleIcon(profile.role)}
                <Badge className={getRoleColor(profile.role)}>
                  {profile.role.charAt(0).toUpperCase() + profile.role.slice(1)}
                </Badge>
              </div>
            </div>
          </div>
          <Button asChild>
            <Link href="/dashboard/profiles/update">
              <Edit className="w-4 h-4 mr-2" />
              Edit Profile
            </Link>
          </Button>
        </CardHeader>
      </Card>

      {/* Contact Information */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center space-x-2">
            <User className="w-5 h-5" />
            <span>Contact Information</span>
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex items-center space-x-3">
            <Mail className="w-4 h-4 text-muted-foreground" />
            <span className="text-sm text-muted-foreground">Email:</span>
            <span>{profile.email}</span>
          </div>
          {profile.phone && (
            <div className="flex items-center space-x-3">
              <Phone className="w-4 h-4 text-muted-foreground" />
              <span className="text-sm text-muted-foreground">Phone:</span>
              <span>{profile.phone}</span>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Biography */}
      {profile.bio && (
        <Card>
          <CardHeader>
            <CardTitle>About</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm leading-relaxed">{profile.bio}</p>
          </CardContent>
        </Card>
      )}

      {/* Role-specific information */}
      <Card>
        <CardHeader>
          <CardTitle>
            {profile.role === 'teacher'
              ? 'Teaching Information'
              : 'Academic Information'}
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          {profile.role === 'student' && (
            <>
              <div className="flex items-center space-x-3">
                <GraduationCap className="w-4 h-4 text-muted-foreground" />
                <span className="text-sm text-muted-foreground">
                  Grade Level:
                </span>
                <span>{profile.grade_level}</span>
              </div>
              {profile.major && (
                <div className="flex items-center space-x-3">
                  <Briefcase className="w-4 h-4 text-muted-foreground" />
                  <span className="text-sm text-muted-foreground">Major:</span>
                  <span>{profile.major}</span>
                </div>
              )}
            </>
          )}
          {profile.role === 'teacher' && (
            <div className="flex items-center space-x-3">
              <Briefcase className="w-4 h-4 text-muted-foreground" />
              <span className="text-sm text-muted-foreground">Position:</span>
              <span>{profile.position}</span>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Account Information */}
      <Card>
        <CardHeader>
          <CardTitle>Account Information</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex items-center space-x-3">
            <span className="text-sm text-muted-foreground">User ID:</span>
            <span className="font-mono text-sm">{profile.user_id}</span>
          </div>
          <div className="flex items-center space-x-3">
            <span className="text-sm text-muted-foreground">Member since:</span>
            <span>{formatDate(Math.floor(profile.created_at / 1000))}</span>
          </div>
          <div className="flex items-center space-x-3">
            <span className="text-sm text-muted-foreground">Last updated:</span>
            <span>{formatDate(Math.floor(profile.updated_at / 1000))}</span>
          </div>
          <div className="flex items-center space-x-3">
            <span className="text-sm text-muted-foreground">Status:</span>
            <Badge variant={profile.is_active ? 'default' : 'secondary'}>
              {profile.is_active ? 'Active' : 'Inactive'}
            </Badge>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
