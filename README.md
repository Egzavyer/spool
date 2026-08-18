# Spool

## Overview

Server receives HTTP request from client to perform a job and validates it. A valid job request is store in the database with a PENDING status. Workers claim jobs and complete them. The system is fault-tolerant and resistant. Workers can complete multiple tasks in parallel.
