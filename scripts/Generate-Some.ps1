$difficulties = @('easy', 'medium', 'hard')
$count = 10
$duration = 20
$size = 8

$difficulties | % {
	$difficulty = $_
	(1..$count) | % {
		go run ./cmd/dragons generate `
		-duration $duration `
		-difficulty $difficulty `
		-w $size -h $size
	}
}
