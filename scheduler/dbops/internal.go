package dbops

import "log"

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmt, err := DBConn.Prepare("select video_id from video_del_rec limit ?")
	if err != nil {
		log.Printf("ReadVideoDeletionRecord error: %v", err)
		return nil, err
	}
	defer stmt.Close()
	var ids []string
	rows, err := stmt.Query(count)
	if err != nil {
		log.Printf("ReadVideoDeletionRecord error: %v", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		if err = rows.Scan(&id); err != nil {
			return ids, nil
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmt, err := DBConn.Prepare("delete from video_del_rec where video_id = ?")
	if err != nil {
		log.Printf("DelVideoDeletionRecord error: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(vid)
	if err != nil {
		log.Printf("DelVideoDeletionRecord error: %v", err)
		return err
	}
	return nil
}

func AddVideoDeletionRecord(vid string) error {
	stmt, err := DBConn.Prepare("insert into video_del_rec(video_id) values (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error : %v", err)
		return err
	}
	return nil
}
