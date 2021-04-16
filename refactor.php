<?php

function findFirstStringInBracket($str)
{
    $firstbracket = strstr($str, '(');
    if (strlen($str) == 0 && !strstr($str, '(')) return '';

    if ($firstbracket && strstr($str, ')')) {
        $firstbracket = ltrim($firstbracket, '(');
        return strstr($firstbracket, ')', true);
    }

    return '';
}