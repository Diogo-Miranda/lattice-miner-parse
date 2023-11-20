library(foreign)
path <- "C:\\Users\\diogo\\Desktop\\dava.sav"
data <- foreign::read.spss(path, to.data.frame = TRUE)
head(data)
