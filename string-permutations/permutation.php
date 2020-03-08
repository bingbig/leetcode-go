<?php

/**
 * 输入：一个字符串数组，如["hello", "hello", "world", "php", "is the", "best language", "in the"]
 * 输出：所有可能的组合，不考虑组合内字符串的排序顺序，如["hello", "world"]和 ["world"，"hello"]视为同一个组合
 */

function preCombine($strArr){
    $strArr = array_flip(array_flip($strArr));
    asort($strArr);
    $strArr = array_values($strArr);
    return $strArr;
}

function combine($input, $m)
{
    if($m == 0){
        return [];
    }

    if($m == 1){
        return array_map(function($item){
            return [$item];
        }, $input);
    }
    $result = [];
    for($i=0;$i<=count($input)-$m;$i++){
        $nextinput=array_slice($input,$i+1);
        $nextresult=combine($nextinput,$m-1);
        foreach($nextresult as $one){
            $result[]=array_merge([$input[$i]], $one);
        }
    }
    return $result;
}

function perm($input, $m)
{
    if($m == 0){
        return [];
    }

    if($m == 1){
        return array_map(function($item){
            return [$item];
        }, $input);
    }

    $result = [];
    for($i=0;$i<count($input);$i++){
        $nextinput=array_merge(array_slice($input,0,$i),array_slice($input,$i+1));
        $nextresult=perm($nextinput,$m-1);
        foreach($nextresult as $one)
        {
            $result[] = array_merge([$input[$i]],$one);
        }
    }
    return $result;
}

// 所有的排列
function allPerm($input)
{
    $result = [];
    for($i = 1;$i <= count($input);$i++){
        $result[] = perm($input, $i);
    }

    return $result;
}

// 所有的组合
function allCombine($input)
{
    $result = [];
    for($i = 1;$i <= count($input);$i++){
        $result[] = combine($input, $i);
    }

    return $result;
}

$strArr = ["1", "2", "3", "4"];

$strArr = preCombine($strArr);
echo json_encode(allCombine($strArr));
echo "\n\n";
echo json_encode(allPerm($strArr));
