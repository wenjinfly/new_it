

//静态方法，获得随机数

export const getRandomNumber = (min, max) => {
    parseInt(Math.random() * (max - min)) + min
}