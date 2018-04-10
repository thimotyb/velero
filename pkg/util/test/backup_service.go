/*
Copyright 2017 the Heptio Ark contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by mockery v1.0.0
package test

import io "io"
import mock "github.com/stretchr/testify/mock"
import time "time"
import v1 "github.com/heptio/ark/pkg/apis/ark/v1"

// BackupService is an autogenerated mock type for the BackupService type
type BackupService struct {
	mock.Mock
}

// CreateSignedURL provides a mock function with given fields: target, bucket, ttl
func (_m *BackupService) CreateSignedURL(target v1.DownloadTarget, bucket, directory string, ttl time.Duration) (string, error) {
	ret := _m.Called(target, bucket, directory, ttl)

	var r0 string
	if rf, ok := ret.Get(0).(func(v1.DownloadTarget, string, string, time.Duration) string); ok {
		r0 = rf(target, bucket, directory, ttl)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v1.DownloadTarget, string, string, time.Duration) error); ok {
		r1 = rf(target, bucket, directory, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBackupDir provides a mock function with given fields: bucket, backupName
func (_m *BackupService) DeleteBackupDir(bucket string, backupName string) error {
	ret := _m.Called(bucket, backupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(bucket, backupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DownloadBackup provides a mock function with given fields: bucket, name
func (_m *BackupService) DownloadBackup(bucket string, name string) (io.ReadCloser, error) {
	ret := _m.Called(bucket, name)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string) io.ReadCloser); ok {
		r0 = rf(bucket, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bucket, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBackups provides a mock function with given fields: bucket
func (_m *BackupService) GetAllBackups(bucket string) ([]*v1.Backup, error) {
	ret := _m.Called(bucket)

	var r0 []*v1.Backup
	if rf, ok := ret.Get(0).(func(string) []*v1.Backup); ok {
		r0 = rf(bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Backup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackup provides a mock function with given fields: bucket, name
func (_m *BackupService) GetBackup(bucket string, name string) (*v1.Backup, error) {
	ret := _m.Called(bucket, name)

	var r0 *v1.Backup
	if rf, ok := ret.Get(0).(func(string, string) *v1.Backup); ok {
		r0 = rf(bucket, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Backup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bucket, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadBackup provides a mock function with given fields: bucket, name, metadata, backup, log
func (_m *BackupService) UploadBackup(bucket string, name string, metadata io.Reader, backup io.Reader, log io.Reader) error {
	ret := _m.Called(bucket, name, metadata, backup, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader, io.Reader, io.Reader) error); ok {
		r0 = rf(bucket, name, metadata, backup, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadRestoreLog provides a mock function with given fields: bucket, backup, restore, log
func (_m *BackupService) UploadRestoreLog(bucket string, backup string, restore string, log io.Reader) error {
	ret := _m.Called(bucket, backup, restore, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, io.Reader) error); ok {
		r0 = rf(bucket, backup, restore, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadRestoreResults provides a mock function with given fields: bucket, backup, restore, results
func (_m *BackupService) UploadRestoreResults(bucket string, backup string, restore string, results io.Reader) error {
	ret := _m.Called(bucket, backup, restore, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, io.Reader) error); ok {
		r0 = rf(bucket, backup, restore, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
