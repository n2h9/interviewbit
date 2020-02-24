<?php
    /**
    * @input $a -> array of integers
    */
    // return an array of integers //

    function wave($a){
        $l = count($a);
        if ($l <= 1) {
            return $a;
        }
        sort($a);
        for ($i = 1; $i < $l; $i+=2) {
            swap($a, $i, $i-1);
        }
        
        return $a;
    }
    
    function swap(&$a, $i, $j) {
        $tmp = $a[$i];
        $a[$i] = $a[$j];
        $a[$j] = $tmp;
    }
?>