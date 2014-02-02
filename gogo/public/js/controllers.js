var gogo = angular.module('gogo', []);
 
gogo.factory('UserFactory', ['$http', '$q', function ($http, $q) {

    return {
        getUser: function () {
            var deferred = $q.defer();

            $http({
                method: 'GET',
                url: '/api/user/1'
            }).success(function(data, status) {
                deferred.resolve(data)
            }).error(function(data, status) {
                deferred.reject(data)
            });

            return deferred.promise;
        }
    };

}]);

gogo.controller('UserCtrl', ['UserFactory', '$scope', function (UserFactory, $scope) {

    $scope.user = {}

    UserFactory.getUser().then(function(data) {
        // call was successful
        $scope.user = data;
    }, function(data) {
        // call returned an error
    });

}]);