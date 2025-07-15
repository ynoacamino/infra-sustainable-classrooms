'use client';

import { useState, useEffect, useRef } from 'react';
import {
  Play,
  Pause,
  Volume2,
  VolumeX,
  Maximize,
  Minimize,
  SkipBack,
  SkipForward,
} from 'lucide-react';
import { Button } from '@/ui/button';
import { Slider } from '@/ui/slider';
import type { VideoDetails } from '@/types/video_learning/models';
import { toast } from 'sonner';

interface VideoPlayerProps {
  video: VideoDetails;
}

export function VideoPlayer({ video }: VideoPlayerProps) {
  const [isPlaying, setIsPlaying] = useState(false);
  const [currentTime, setCurrentTime] = useState(0);
  const [duration, setDuration] = useState(0);
  const [volume, setVolume] = useState(1);
  const [isMuted, setIsMuted] = useState(false);
  const [isFullscreen, setIsFullscreen] = useState(false);
  const [showControls, setShowControls] = useState(true);
  const [playbackRate, setPlaybackRate] = useState(1);
  const [buffered, setBuffered] = useState(0);

  const videoRef = useRef<HTMLVideoElement>(null);
  const playerRef = useRef<HTMLDivElement>(null);
  const controlsTimeoutRef = useRef<NodeJS.Timeout>(null);

  useEffect(() => {
    const videoElement = videoRef.current;
    if (!videoElement) return;

    const handleTimeUpdate = () => {
      setCurrentTime(videoElement.currentTime);
    };

    const handleDurationChange = () => {
      setDuration(videoElement.duration);
    };

    const handlePlay = () => {
      setIsPlaying(true);
    };

    const handlePause = () => {
      setIsPlaying(false);
    };

    const handleVolumeChange = () => {
      setVolume(videoElement.volume);
      setIsMuted(videoElement.muted);
    };

    const handleProgress = () => {
      if (videoElement.buffered.length > 0) {
        const bufferedEnd = videoElement.buffered.end(
          videoElement.buffered.length - 1,
        );
        setBuffered((bufferedEnd / duration) * 100);
      }
    };

    const handleFullscreenChange = () => {
      setIsFullscreen(!!document.fullscreenElement);
    };

    videoElement.addEventListener('timeupdate', handleTimeUpdate);
    videoElement.addEventListener('durationchange', handleDurationChange);
    videoElement.addEventListener('play', handlePlay);
    videoElement.addEventListener('pause', handlePause);
    videoElement.addEventListener('volumechange', handleVolumeChange);
    videoElement.addEventListener('progress', handleProgress);
    document.addEventListener('fullscreenchange', handleFullscreenChange);

    return () => {
      videoElement.removeEventListener('timeupdate', handleTimeUpdate);
      videoElement.removeEventListener('durationchange', handleDurationChange);
      videoElement.removeEventListener('play', handlePlay);
      videoElement.removeEventListener('pause', handlePause);
      videoElement.removeEventListener('volumechange', handleVolumeChange);
      videoElement.removeEventListener('progress', handleProgress);
      document.removeEventListener('fullscreenchange', handleFullscreenChange);
    };
  }, [duration]);

  const togglePlayPause = () => {
    const videoElement = videoRef.current;
    if (!videoElement) return;

    if (isPlaying) {
      videoElement.pause();
    } else {
      videoElement.play();
    }
  };

  const handleSeek = (value: number[]) => {
    const videoElement = videoRef.current;
    if (!videoElement) return;

    const newTime = (value[0] / 100) * duration;
    videoElement.currentTime = newTime;
    setCurrentTime(newTime);
  };

  const handleVolumeChange = (value: number[]) => {
    const videoElement = videoRef.current;
    if (!videoElement) return;

    const newVolume = value[0] / 100;
    videoElement.volume = newVolume;
    setVolume(newVolume);
    setIsMuted(newVolume === 0);
  };

  const toggleMute = () => {
    const videoElement = videoRef.current;
    if (!videoElement) return;

    videoElement.muted = !videoElement.muted;
    setIsMuted(!isMuted);
  };

  const toggleFullscreen = async () => {
    const playerElement = playerRef.current;
    if (!playerElement) return;

    try {
      if (!document.fullscreenElement) {
        await playerElement.requestFullscreen();
      } else {
        await document.exitFullscreen();
      }
    } catch (error) {
      console.error('Fullscreen toggle failed:', error);
    }
  };

  const skip = (seconds: number) => {
    const videoElement = videoRef.current;
    if (!videoElement) return;

    videoElement.currentTime += seconds;
  };

  const changePlaybackRate = (rate: number) => {
    const videoElement = videoRef.current;
    if (!videoElement) return;

    videoElement.playbackRate = rate;
    setPlaybackRate(rate);
  };

  const formatTime = (time: number) => {
    if (isNaN(time)) return '0:00';

    const minutes = Math.floor(time / 60);
    const seconds = Math.floor(time % 60);
    return `${minutes}:${seconds.toString().padStart(2, '0')}`;
  };

  const handleMouseMove = () => {
    setShowControls(true);

    if (controlsTimeoutRef.current) {
      clearTimeout(controlsTimeoutRef.current);
    }

    controlsTimeoutRef.current = setTimeout(() => {
      if (isPlaying) {
        setShowControls(false);
      }
    }, 3000);
  };

  const handleMouseLeave = () => {
    if (isPlaying) {
      setShowControls(false);
    }
  };

  return (
    <div
      ref={playerRef}
      className={`relative bg-black rounded-lg overflow-hidden ${
        isFullscreen ? 'fixed inset-0 z-50 rounded-none' : 'aspect-video'
      }`}
      onMouseMove={handleMouseMove}
      onMouseLeave={handleMouseLeave}
    >
      {/* Video Element */}
      <video
        ref={videoRef}
        src={video.video_url}
        className="w-full h-full object-contain"
        onClick={togglePlayPause}
        onError={() => toast.error('Failed to load video file')}
      />

      {/* Controls Overlay */}
      <div
        className={`absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/80 to-transparent p-4 transition-opacity duration-300 ${
          showControls ? 'opacity-100' : 'opacity-0'
        }`}
      >
        {/* Progress Bar */}
        <div className="mb-4">
          <div className="relative">
            {/* Buffered Progress */}
            <div
              className="absolute top-0 left-0 h-1 bg-white/30 rounded-full"
              style={{ width: `${buffered}%` }}
            />
            {/* Seek Bar */}
            <Slider
              value={[duration ? (currentTime / duration) * 100 : 0]}
              onValueChange={handleSeek}
              className="w-full"
              step={0.1}
            />
          </div>
        </div>

        {/* Controls */}
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Button
              variant="ghost"
              size="icon"
              onClick={togglePlayPause}
              className="text-white hover:bg-white/20"
            >
              {isPlaying ? (
                <Pause className="h-5 w-5" />
              ) : (
                <Play className="h-5 w-5" />
              )}
            </Button>

            <Button
              variant="ghost"
              size="icon"
              onClick={() => skip(-10)}
              className="text-white hover:bg-white/20"
            >
              <SkipBack className="h-4 w-4" />
            </Button>

            <Button
              variant="ghost"
              size="icon"
              onClick={() => skip(10)}
              className="text-white hover:bg-white/20"
            >
              <SkipForward className="h-4 w-4" />
            </Button>

            <div className="flex items-center gap-2 ml-4">
              <Button
                variant="ghost"
                size="icon"
                onClick={toggleMute}
                className="text-white hover:bg-white/20"
              >
                {isMuted || volume === 0 ? (
                  <VolumeX className="h-4 w-4" />
                ) : (
                  <Volume2 className="h-4 w-4" />
                )}
              </Button>

              <div className="w-20">
                <Slider
                  value={[isMuted ? 0 : volume * 100]}
                  onValueChange={handleVolumeChange}
                  className="w-full"
                />
              </div>
            </div>

            <div className="text-white text-sm ml-4">
              {formatTime(currentTime)} / {formatTime(duration)}
            </div>
          </div>

          <div className="flex items-center gap-2">
            <select
              value={playbackRate}
              onChange={(e) => changePlaybackRate(parseFloat(e.target.value))}
              className="bg-transparent text-white text-sm border border-white/20 rounded px-2 py-1"
            >
              <option value={0.5} className="text-black">
                0.5x
              </option>
              <option value={0.75} className="text-black">
                0.75x
              </option>
              <option value={1} className="text-black">
                1x
              </option>
              <option value={1.25} className="text-black">
                1.25x
              </option>
              <option value={1.5} className="text-black">
                1.5x
              </option>
              <option value={2} className="text-black">
                2x
              </option>
            </select>

            <Button
              variant="ghost"
              size="icon"
              onClick={toggleFullscreen}
              className="text-white hover:bg-white/20"
            >
              {isFullscreen ? (
                <Minimize className="h-4 w-4" />
              ) : (
                <Maximize className="h-4 w-4" />
              )}
            </Button>
          </div>
        </div>
      </div>

      {/* Center Play Button (when paused) */}
      {!isPlaying && (
        <div className="absolute inset-0 flex items-center justify-center">
          <Button
            variant="ghost"
            size="icon"
            onClick={togglePlayPause}
            className="text-white hover:bg-white/20 w-16 h-16 rounded-full"
          >
            <Play className="h-8 w-8" />
          </Button>
        </div>
      )}
    </div>
  );
}
