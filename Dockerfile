## ogwc (https://github.com/c-mueller/ogwc).
## Copyright (C) 2018-2020 Christian Müller <dev@c-mueller.xyz>.
##
## This program is free software: you can redistribute it and/or modify
## it under the terms of the GNU Affero General Public License as published by
## the Free Software Foundation, either version 3 of the License, or
## (at your option) any later version.
##
## This program is distributed in the hope that it will be useful,
## but WITHOUT ANY WARRANTY; without even the implied warranty of
## MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
## GNU Affero General Public License for more details.
##
## You should have received a copy of the GNU Affero General Public License
## along with this program.  If not, see <https://www.gnu.org/licenses/>.

FROM ubuntu:18.04 AS runtime
WORKDIR /ogwc
COPY build/ogwc /bin/ogwc
EXPOSE 16666
ENV REDIS_ADDR=myredis:6379
ENV REDIS_DB=0
ENV REDIS_PW=""
ENV BIND_ADDR=":16666"
CMD ogwc server --bind-address="$BIND_ADDR" --redis-address="$REDIS_ADDR" --redis-database="$REDIS_DB" --redis-password="$REDIS_PASSWORD"