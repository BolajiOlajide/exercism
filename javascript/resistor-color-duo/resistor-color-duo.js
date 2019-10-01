//
// This is only a SKELETON file for the 'Resistor Color Duo' exercise. It's been provided as a
// convenience to get you started writing code faster.
//
import { colorObject } from '../resistor-color/resistor-color';

export const value = (colors) => {
    const [colorOne, colorTwo] = colors;

    return parseInt(`${colorObject[colorOne]}${colorObject[colorTwo]}`, 10);
};
