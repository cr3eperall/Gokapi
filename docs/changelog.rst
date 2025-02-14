.. _changelog:


Changelog
=========

Overview of all Changes
-----------------------

v1.8.1: 7 Feb 2024
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* Reworked OIDC authentication, added support for Groups, use consent instead of selec_account, better error handling
* Added wildcard support for OIDC groups and users
* Fixed crash on client timeout #125
* Added /auth/create API endpoint for creating API keys
* Minor changes and fixes


v1.8.0: 9 Dec 2023
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Changed Database to Sqlite3
* Dropped Windows 32bit support
* Only 4,000 parallel requests that are writing to the database are supported now, any requests above that limit may be rejected. Up to 500,000 parallel reading requests were tested.
* According to the documentation, the GOKAPI_DATA_DIR environment variable should be persistent, however that was not the case. Now the data directory that was set on first start will be used. If you were using GOKAPI_DATA_DIR after the first start, make sure that the data directory is the one found in your config file.
* By default, IP addresses of clients downloading files are not saved anymore to comply with GDPR. This can be enabled by re-running the setup
* Existing API keys will be granted all API permissions except MODIFY_API, therefore cannot use /auth/friendlyname without having the permission granted first
* The undocumented GOKAPI_FILE_DB environment variable was removed
* Removed optional application for reading database content
* Parameters of already uploaded files can be edited now
* Added permission model for API tokens
* Added /auth/modify and /files/modify API endpoint
* Fixed "Powered by Gokapi" URL not clickable
* Fixed the ASCII logo #108 Thanks to @Kwonunn
* Improved UI
* Fixed minor bugs
* Updated dependencies
* Updated documentation


v1.7.2: 13 May 2023
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Added option to change the name in the setup
* The filename is now shown in the title for downloads
* SessionStorage is used instead of localStorage for e2e decryption
* Replaced expiry image with dynamic SVG


v1.7.1: 14 Apr 2023
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Fixed Gokapi not able to upload when running on a Windows system #95
* Improved Upload UI
* Added healthcheck for docker by @Jisagi in #89
* Fixed upload counter not updating after upload #92
* Fixed hotlink generation on files that required client-side decryption
* Replaced go:generate code with native Go
* Min Go version now 1.20
* Updated dependencies
* A lot of refactoring, minor changes
* Fixed background not loading in 1.7.0 (unpublished release) #101

v1.6.2: 14 Feb 2023
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Fixed timeout if a large file was uploaded to the cloud #81
* File overview is now sortable and searchable
* Added log viewer
* Updated Go to 1.20
* Other minor changes and fixes

v1.6.1: 17 Aug 2022
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Fixed setup throwing error 500 on docker installation


v1.6.0: 17 Aug 2022
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Use chunked uploads instead of single upload #68
* Add end-to-end encryption #71
* Fixed hotlink not being generated for uploads through API with unlimited storage time
* Added arm64 to Docker latest image
* Added API call to duplicate existing files
* Fixed bug where encrypted files could not be downloaded after rerunning setup
* Port selection is now disabled when running setup with docker
* Added timeout for AWS if endpoint is invalid
* Added flag to disable CORS check on startup
* Service worker for insecure connections is now hosted on Github
* "Noaws" version is not included as binary build anymore, but can be generated manually


v1.5.2: 08 Jun 2022
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Added ARMv8 (ARM64) to Docker image
* Added option to always store images locally in order to support hotlink for encrypted files
* Fixed crash when remote files exist but system was changed to local files after running --reconfigure
* Added warning if incorrect CORS setting are set for AWS bucket
* Added button in setup to test AWS credentials
* Added more build infos to --version output
* Added download counter
* Added flags for port, config and data location, better flag usage overview
* Fixed that a file was reuploaded to AWS, even if it already existed
* Fixed error image for hotlinks not displaying if nosniff is enforced
* Fixed that two text files were created when pasting text
* Fixed docker image in documentation @emanuelduss

v1.5.1: 10 Mar 2022
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Fixed that selection of remote storage was not available during intitial setup
* Fixed that "bind to localhost" could be selected on docker image during initial setup
* Fixed that with Level 1 encryption remote files were encrypted as well
* If Gokapi is hosted under a https URL, the serviceworker for remote decryption is now included, which fixes that Firefox users with restrictive settings could not download encrypted files from remote storage
* Design improvements by @mraif13


v1.5.0: 08 Mar 2022
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Minimum version for upgrading is 1.3
* Encryption support for local and remote files
* Additional authentication methods: Header-Auth, OIDC and Reverse Proxy
* Option to allow unlimited downloads of files
* The configuration file has been partly replaced with a database. After the first start, the configuration file may be read-only
* A web-based setup instead of command line


v1.3.1: 03 Jul 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
* Default upload limit is now 100GB and can be changed with environment variables on first start
* Fixed upload not working when using suburl on webserver for Gokapi
* Added log file
* Minor performance increase

v1.3.0: 17 May 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* Added cloudstorage support (AWS S3 / Backblaze B2)
* After changing password, all sessions will be logged out
* Fixed terminal input on Windows
* Added SSL support
* Documentation now hosted on ReadTheDocs

v1.2.0: 07 May 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* Fixed Docker images
* Added API
* Added header to prevent caching by browser / proxy
* Fixed upload timeout
* Added timeouts for server
* Added header to show download progress
* Prevent data races
* Cleanup routine does not delete files anymore while they are being downloaded
* Fixed that env ``LENGTH_ID`` was being ignored
* Show message if docker container is run on initial setup without ``-it``
* A lot of refactoring and minor improvements / bug fixes

v1.1.3: 07 Apr 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* Fixed bug where salts were not used anymore for password hashing
* Added hotlinking for image files
* Added logout button

v1.1.2: 03 Apr 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* Added support for env variables, major refactoring
* Configurations like length of the ID or salts can be changed with env variables now
* Fixed minor bugs, minor enhancements

v1.1.0: 18 Mar 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* Added option to password protect uploads
* Added ability to paste images into admin upload


v1.0.1: 12 Mar 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* Increased security of generated download IDs


v1.0: 12 Mar 2021
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

* First stable release of the program


Upgrading
-----------------------

Upgrading to 1.8
^^^^^^^^^^^^^^^^^^

* You need to update to Gokapi 1.7 before updating to Gokapi 1.8
* With this release, the old key-value database was changed to sqlite3. Please backup all Gokapi data before installing this release. On first start, the old database will be migrated and all users will be logged out. 

Upgrading to 1.5
^^^^^^^^^^^^^^^^^^

* You need to update to Gokapi 1.3 before updating to Gokapi 1.5
* After the upgrade the config file can be read-only
* Initial setup has to be done through a web interface now, setting Gokapi up through env variables is not possible anymore
* If you would like to use new features like a different authentication method, please run Gokapi with the paramter ``--reconfigure`` to open the setup  
* If you set the length of the file ID to 80 or more, you need to delete all files before running this update

Upgrading to 1.3
^^^^^^^^^^^^^^^^^^

* If you would like to use native SSL, please pass the environment variable ``GOKAPI_USE_SSL`` on first start after the update or manually edit the configuration file
* AWS S3 and Backblaze B2 can now be used instead of local storage! Please refer to the documentation on how to set it up.
