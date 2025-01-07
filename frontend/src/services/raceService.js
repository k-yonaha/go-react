const API_URL = "http://localhost:8080/api/race-schedule";  // レース情報のエンドポイント

// 部屋名（course_name）を基に次のレースを取得
export const getNextRaceByCourse = async (courseName) => {
  try {
    const response = await fetch(`${API_URL}?course_name=${courseName}`);
    if (!response.ok) {
      throw new Error("次のレース情報の取得に失敗しました");
    }
    return await response.json();
  } catch (error) {
    throw new Error(error.message);
  }
};

export const getRaceSchedulesByDate = async () => {
    try {
      const response = await fetch(`${API_URL}/today`);
      if (!response.ok) {
        throw new Error("レース情報の取得に失敗しました");
      }
      const data = await response.json();

      // レース情報を course_name ごとにグループ化
    //   const raceSchedulesByCourse = data.reduce((acc, raceSchedule) => {
    //     const { course_name } = raceSchedule;
    //     if (!acc[course_name]) {
    //       acc[course_name] = [];
    //     }
    //     acc[course_name].push(raceSchedule);
    //     return acc;
    //   }, {});
      
      return data;
    } catch (error) {
      throw new Error(error.message);
    }
  };