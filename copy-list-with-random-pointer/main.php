
<?php

class Node {
    public $val;
    public $next;
    public $random;

    // @param Integer $val
    // @param Node $next
    // @param Node $random
    function __construct($val, $next, $random) {
        $this->val = $val;
        $this->next = $next;
        $this->random = $random;
    }
}

class Solution {

    /**
     * @param Node $head
     * @return Node
     */
    function copyRandomList($head) {
        if($head == null){
            return null;
        }

        $p = $head;
        while($p != null){
            $cnode = new Node($p->val, $p->next, null);
            $p->next = $cnode;
            $p = $cnode->next;
        }

        $p = $head;
        while ($p != null) {
            if($p->random != null) {
                $p->next->random = $p->random->next;
            }
            $p = $p->next->next;
        }

        $ori = $head;
        $new = $head->next;
        $p1 = $ori;
        $p2 = $new;
        $p = $head->next->next;
        while($p != null){
            $p1->next = $p;
            $p2->next = $p->next;
            $p1 = $p1->next;
            $p2 = $p2->next;
            $p = $p->next->next;
        }

        $head = $ori;
        return $new;
    }

    function printRandomList($head) {
        $p = $head;
        while($p!=null){
            printf("loc:%s ,val: %d, next: %s, random: %s\n", spl_object_hash($p), $p->val, $p->next->val ?? 'null', $p->random->val ?? null);
            $p = $p->next;
        }
    }
}

$solution = New Solution();
$head = new Node(7,null,null);
$node2 = new Node(13, null, $head);
$node3 = new Node(11, null, null);
$node4 = new Node(10, null , $node2);
$node5 = new Node(1, null, $head);

$head->next = $node2;
$node2->next = $node3;
$node3->next = $node4;
$node4->next = $node5;
$node3->random = $node4;

$solution->printRandomList($head);

echo "\n\n";

$cloned = $solution->copyRandomList($head);
$solution->printRandomList($cloned);

