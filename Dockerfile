FROM ubuntu
WORKDIR /app
COPY ./build /app
EXPOSE 8000
CMD /app/lp_ss_linux