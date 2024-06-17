part 'space_values.dart';

// dammy
class Space {
  final double x;
  final double y;
  final dynamic key;

  const Space(double space, {this.key})
      : x = space,
        y = space;

  const Space.y(double space, {this.key})
      : x = 0,
        y = space;

  const Space.x(double space, {this.key})
      : x = space,
        y = 0;
}
