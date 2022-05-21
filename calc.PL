my $num1 = 4;
my $num2 = 2;

sub calc{
    my ($num1, $num2) = @_;
    my $result = $num1 + $num2;
    print "$num1 + $num2 = $result\n";
}

sub div{
    my ($num1, $num2) = @_;
    my $result = $num1 / $num2;
    print "$num1 / $num2 = $result\n";
}

sub multi{
    my ($num1, $num2) = @_;
    my $result = $num1 * $num2;
    print "$num1 * $num2 = $result\n";
}

sub minus{
    my ($num1, $num2) = @_;
    my $result = $num1 - $num2;
    print "$num1 - $num2 = $result";
}

calc($num1, $num2);
div($num1, $num2);
multi($num1, $num2);
minus($num1, $num2);
