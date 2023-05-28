INSERT INTO `questions` VALUES (1, '2023-4-20 23:49:58', '2023-4-20 23:49:58', NULL, '斐波那契数列', '写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：\n\nF(0) = 0,   F(1) = 1\nF(N) = F(N - 1) + F(N - 2), 其中 N > 1.\n斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。\n');
INSERT INTO `questions` VALUES (6, '2023-4-22 00:49:18', '2023-4-22 00:49:18', NULL, '快速排序', '实现一个快速排序算法，要求能够对一个整数数组进行排序');
INSERT INTO `questions` VALUES (7, '2023-4-22 05:20:29', '2023-4-22 05:20:29', NULL, '堆排序', '实现一个堆排序算法，要求能够对一个整数数组进行排序');
INSERT INTO `questions` VALUES (8, '2023-4-22 05:26:22', '2023-4-22 05:26:22', NULL, '冒泡排序', '实现一个冒泡排序算法，要求能够对一个整数数组进行排序');
INSERT INTO `questions` VALUES (9, '2023-4-22 07:20:47', '2023-4-22 07:20:47', NULL, '归并排序', '实现一个归并排序算法，要求能够对一个整数数组进行排序');
INSERT INTO `questions` VALUES (10, '2023-4-22 08:25:13', '2023-4-22 08:25:13', NULL, '数据流中的中位数', '如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。');
INSERT INTO `questions` VALUES (11, '2023-4-22 09:02:37', '2023-4-22 09:02:37', NULL, '两数之和', '给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。\n\n你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。\n');
INSERT INTO `questions` VALUES (12, '2023-4-22 09:08:31', '2023-4-22 09:08:31', NULL, '正则表达式匹配', '给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 \'.\' 和 \'*\' 的正则表达式匹配。\n\n\'.\' 匹配任意单个字符\n\'*\' 匹配零个或多个前面的那一个元素\n所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。');
INSERT INTO `questions` VALUES (13, '2023-4-22 09:15:52', '2023-4-22 09:15:52', NULL, '搜索插入位置', '给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。');
INSERT INTO `questions` VALUES (14, '2023-4-22 09:55:14', '2023-4-22 09:55:14', NULL, '跳跃游戏 II', '给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。\n\n每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:\n\n0 <= j <= nums[i] \ni + j < n\n返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。\n\n');
INSERT INTO `questions` VALUES (15, '2023-4-22 10:01:08', '2023-4-22 10:01:08', NULL, '跳跃游戏 ', '给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。\n\n数组中的每个元素代表你在该位置可以跳跃的最大长度。\n\n判断你是否能够到达最后一个下标。\n\n ');
INSERT INTO `questions` VALUES (16, '2023-4-22 10:04:27', '2023-4-22 10:04:27', NULL, '股票的最大利润', '假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？\n\n ');
INSERT INTO `questions` VALUES (17, '2023-4-22 10:14:35', '2023-4-22 10:14:35', NULL, '买卖股票的最佳时机', '给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。\n\n你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。\n\n返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。\n\n');
INSERT INTO `questions` VALUES (18, '2023-4-22 10:33:51', '2023-4-22 10:33:51', NULL, '买卖股票的最佳时机 II', '给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。\n\n在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。\n\n返回 你能获得的 最大 利润 。\n');
INSERT INTO `questions` VALUES (19, '2023-4-22 10:35:48', '2023-4-22 10:35:48', NULL, '买卖股票的最佳时机 III', '给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。\n\n设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。\n\n注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。\n');
INSERT INTO `questions` VALUES (20, '2023-4-22 10:40:06', '2023-4-22 10:40:06', NULL, '最佳买卖股票时机含冷冻期', '给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​\n\n设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:\n\n卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。\n注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。\n');
INSERT INTO `questions` VALUES (21, '2023-4-22 10:45:15', '2023-4-22 10:45:15', NULL, '无重复字符的最长子串', '给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。');
INSERT INTO `questions` VALUES (22, '2023-4-22 10:52:17', '2023-4-22 10:52:17', NULL, '位运算实现加法', '位运算实现加法');
INSERT INTO `questions` VALUES (23, '2023-4-22 10:56:37', '2023-4-22 10:56:37', NULL, '位运算实现乘法', '位运算实现乘法');
INSERT INTO `questions` VALUES (24, '2023-4-22 10:58:51', '2023-4-22 10:58:51', NULL, '位运算实现除法', '位运算实现除法');
INSERT INTO `questions` VALUES (25, '2023-4-22 11:04:08', '2023-4-22 11:04:08', NULL, '位运算实现减法', '位运算实现减法');
INSERT INTO `questions` VALUES (26, '2023-4-22 11:05:52', '2023-4-22 11:05:52', NULL, '加一', '给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。\n\n最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。\n\n你可以假设除了整数 0 之外，这个整数不会以零开头。\n');
INSERT INTO `questions` VALUES (27, '2023-4-22 11:09:35', '2023-4-22 11:09:35', NULL, '统计只差一个字符的子串数目', '给你两个字符串 s 和 t ，请你找出 s 中的非空子串的数目，这些子串满足替换 一个不同字符 以后，是 t 串的子串。换言之，请你找到 s 和 t 串中 恰好 只有一个字符不同的子字符串对的数目。\n\n比方说， \"computer\" and \"computation\" 只有一个字符不同： \'e\'/\'a\' ，所以这一对子字符串会给答案加 1 。\n\n请你返回满足上述条件的不同子字符串对数目。\n\n一个 子字符串 是一个字符串中连续的字符。\n');
INSERT INTO `questions` VALUES (28, '2023-4-22 11:11:48', '2023-4-22 11:11:48', NULL, '两数相除', '给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。\n\n整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。\n\n返回被除数 dividend 除以除数 divisor 得到的 商 。\n\n注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−231,  231 − 1] 。本题中，如果商 严格大于 231 − 1 ，则返回 231 − 1 ；如果商 严格小于 -231 ，则返回 -231 。\n');
INSERT INTO `questions` VALUES (29, '2023-4-22 11:14:44', '2023-4-22 11:14:44', NULL, '盛最多水的容器', '给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。\n\n找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。\n\n返回容器可以储存的最大水量。\n\n说明：你不能倾斜容器。\n');
INSERT INTO `questions` VALUES (30, '2023-4-22 11:17:11', '2023-4-22 11:17:11', NULL, '接雨水', '给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。');
INSERT INTO `questions` VALUES (31, '2023-4-23 13:28:14', '2023-4-23 13:28:14', NULL, '字符串相乘', '给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。\n\n注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。\n');
INSERT INTO `questions` VALUES (32, '2023-4-23 13:32:46', '2023-4-23 13:32:46', NULL, '括号生成', '数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。');
INSERT INTO `questions` VALUES (33, '2023-4-23 13:34:13', '2023-4-23 13:34:13', NULL, 'Pow(x, n)', '实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。');