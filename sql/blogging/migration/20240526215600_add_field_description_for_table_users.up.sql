ALTER TABLE users
    ADD COLUMN description text NULL AFTER email;

SET @descriptions = '[
    "Một người yêu thích công nghệ và luôn tìm kiếm những thách thức mới trong lập trình.",
    "Đam mê sáng tạo nội dung và luôn khao khát truyền cảm hứng cho mọi người.",
    "Yêu thích du lịch và khám phá các nền văn hóa khác nhau.",
    "Một người đam mê nghệ thuật và dành thời gian rảnh để vẽ tranh.",
    "Luôn nhiệt huyết với công việc và mong muốn đóng góp tích cực cho cộng đồng.",
    "Đam mê đọc sách và luôn tìm kiếm tri thức mới mỗi ngày.",
    "Một người yêu thể thao và thường xuyên tham gia các hoạt động ngoài trời.",
    "Yêu thích nấu ăn và luôn tìm kiếm những công thức mới để thử nghiệm.",
    "Đam mê âm nhạc và đang học chơi guitar.",
    "Một người đam mê viết lách và mong muốn trở thành một nhà văn.",
    "Yêu thích công việc tư vấn tâm lý và mong muốn giúp đỡ người khác.",
    "Luôn nỗ lực học hỏi để trở thành một chuyên gia trong lĩnh vực của mình.",
    "Đam mê thiết kế thời trang và yêu thích phong cách cổ điển.",
    "Một người đam mê nhiếp ảnh và luôn tìm kiếm khoảnh khắc đẹp trong cuộc sống.",
    "Yêu thích đọc sách khoa học và luôn tò mò về thế giới xung quanh.",
    "Một người luôn hướng đến sự hoàn thiện bản thân qua từng ngày.",
    "Đam mê tổ chức sự kiện và luôn tìm kiếm cơ hội để kết nối mọi người.",
    "Một người yêu thích lập trình web và luôn cập nhật kiến thức mới trong lĩnh vực này.",
    "Luôn sẵn lòng tham gia các hoạt động tình nguyện và giúp đỡ cộng đồng.",
    "Đam mê kinh doanh và luôn tìm kiếm cơ hội khởi nghiệp.",
    "Yêu thích yoga và luôn tìm kiếm sự cân bằng trong cuộc sống.",
    "Một người yêu thích thể thao điện tử và thường xuyên tham gia các giải đấu.",
    "Đam mê công nghệ tài chính và luôn cập nhật xu hướng mới nhất trong ngành.",
    "Yêu thích chăm sóc sức khỏe và thường xuyên tập luyện thể dục.",
    "Một người yêu thích thiên nhiên và thường xuyên tham gia các hoạt động dã ngoại.",
    "Đam mê làm vườn và yêu thích chăm sóc cây cối.",
    "Luôn tìm kiếm cơ hội học hỏi và phát triển kỹ năng mềm.",
    "Đam mê phim ảnh và thường xuyên theo dõi các bộ phim mới.",
    "Yêu thích công việc biên dịch và luôn tìm kiếm cơ hội để cải thiện kỹ năng ngôn ngữ.",
    "Một người luôn đam mê và cống hiến hết mình cho công việc."
]';

SET @cnt = JSON_LENGTH(@descriptions);

UPDATE users
SET description = JSON_UNQUOTE(JSON_EXTRACT(@descriptions, CONCAT('$[', FLOOR(RAND() * @cnt), ']')));
