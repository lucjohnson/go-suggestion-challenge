'use strict'

angular.module('Suggestions', [])
    .controller('SuggController', function($scope, $http) {
        $scope.search = function(prefix, max) {
            if (prefix.trim().length > 0) {
                prefix = prefix.trim()
                $http.get('/api/v1/suggestions?prefix='+ prefix.toLowerCase() + '&max=' + max)
                    .then(function(response) {
                        $scope.Suggestions = response.data.Suggestions;
                        $scope.File = response.data.File;
                        if ($scope.File === './data/wordsEn.txt') {
                            $scope.File = 'dictionary';
                        } else {
                            $scope.File = 'wikipedia'
                        }
                        $scope.Loading = null
                    })
                    .catch(function(err) {
                        $scope.Loading = err.data.Message
                    })
            } else {
                $scope.Suggestions = null;
            }
        }
    })