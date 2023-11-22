library(foreign)

path <- "C:\\Users\\diogo\\Desktop\\data2.sav"
data <- foreign:read.spss(path, to.data.frame = TRUE)

sd <- subset(data, select = c(seq(7, 79)))
sd <- subset(sd, select = -c(SOMS54_RECOD_2017))
print(nrow(sd))

d_zero <- sd[apply(sd, 1, function(row) all(row != 0)), ]
d <- na.omit(d_zero)
df <- data.frame(d)

for (c in colnames(df)) {
  df[[c]] <- ifelse(df[[c]] %in% c(1, 2,3), 1, 0)
}
print(nrow(df))
df <- as.data.frame(df)
head(df)

write.csv(df, file = "C:\\Users\\diogo\\Desktop\\data2.csv")
