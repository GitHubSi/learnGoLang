<?php
/**
 * Created by PhpStorm.
 * User: neojos
 * Date: 2017/5/13
 * Time: 17:37
 */
$n = 41;
$m = 3;
$array = range(1, $n);

do {
    $length = count($array);
    for ($i = 1; $i <= $length; $i++) {
        if ($i % $m == 0) {
            unset($array[$i]);
        }
    }
} while (count($array) > $m);


var_dump($array);
