'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/ui/card';
import { Progress } from '@/ui/progress';
import { Badge } from '@/ui/badge';
import { Trophy, Target, TrendingUp } from 'lucide-react';

export function UserProgress() {
  // Mock user progress data
  const userStats = {
    totalSolved: 5,
    totalExercises: 25,
    successRate: 71,
    currentStreak: 3,
    level: 'Beginner',
    nextLevel: 'Intermediate',
    pointsToNextLevel: 250,
    currentPoints: 180,
  };

  const progressPercentage =
    (userStats.totalSolved / userStats.totalExercises) * 100;
  const levelProgress =
    (userStats.currentPoints / userStats.pointsToNextLevel) * 100;

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Target className="w-5 h-5" />
            Progress Overview
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <div className="flex justify-between text-sm">
              <span>Exercises Completed</span>
              <span>
                {userStats.totalSolved}/{userStats.totalExercises}
              </span>
            </div>
            <Progress value={progressPercentage} className="h-2" />
            <div className="text-xs text-muted-foreground">
              {progressPercentage.toFixed(1)}% completed
            </div>
          </div>

          <div className="flex justify-between items-center">
            <div>
              <div className="text-sm text-muted-foreground">Success Rate</div>
              <div className="text-2xl font-bold text-green-600">
                {userStats.successRate}%
              </div>
            </div>
            <div>
              <div className="text-sm text-muted-foreground">
                Current Streak
              </div>
              <div className="text-2xl font-bold text-orange-600">
                {userStats.currentStreak}
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Trophy className="w-5 h-5" />
            Level Progress
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex items-center gap-2">
            <Badge variant="default">{userStats.level}</Badge>
            <TrendingUp className="w-4 h-4 text-muted-foreground" />
            <Badge variant="outline">{userStats.nextLevel}</Badge>
          </div>

          <div className="space-y-2">
            <div className="flex justify-between text-sm">
              <span>Points to Next Level</span>
              <span>
                {userStats.currentPoints}/{userStats.pointsToNextLevel}
              </span>
            </div>
            <Progress value={levelProgress} className="h-2" />
            <div className="text-xs text-muted-foreground">
              {userStats.pointsToNextLevel - userStats.currentPoints} points
              needed
            </div>
          </div>

          <div className="text-sm text-muted-foreground">
            Keep solving challenges to reach the next level!
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
