import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import "../styles.css"

function Homepage(): JSX.Element {
    return (
        <div>
            <div className="header">
                <br />
                <div className="border">
                    <h3>ยินดีต้อนรับสู่</h3>
                    <h1>ระบบฟาร์มมาร์ท</h1>
                <div className="project-detail">
                ระบบ Farm mart เป็นระบบที่พนักงานสามารถ login เข้าระบบเพื่อคิดเงินสินค้าที่ลูกค้าเลือกซื้อ ภายในร้าน Farm mart โดยที่พนักงานทุกคนสามารถ login เข้าใช้งานระบบผ่าน ID ของพนักงานได้ เป็นระบบที่สามารถจัดการสินค้าภายในร้าน รวมทั้งมีข้อมูลรายละเอียดของสินค้าและ ลูกค้าที่สมัครเป็นสมาชิก                </div>
                </div>
                <br /><br /><br />
                <br /><br /><br />
                <br /><br /><br />
                <br /><br /><br />
                
                <div className="member">สมาชิก</div>
            </div>
            <div className="container">
                <div className="column">
                    <a href="https://www.facebook.com/profile.php?id=100009764586540" className="fa">
                        <img src="https://media.tenor.com/I79_XhbEvhkAAAAC/jiraiya-thumbs-up.gif" alt="เม" width="270" height="260" className="avatar"></img>
                        <Button color="primary" variant="contained">
                            <Typography
                                variant="button"
                            >
                                <div className="devname">นายเมธัส ภาคภูมิพงศ์</div>
                            </Typography>
                        </Button>
                    </a>
                </div>
                <div className="column">
                    <a href="https://www.facebook.com/kengneeha" className="fa">
                        <img src="https://cdn.discordapp.com/attachments/996739649596821574/1034382144010715176/meme-gif-pfp-1.gif" alt="เก่ง" width="270" height="260" className="avatar"></img>
                        <Button color="primary" variant="contained">
                            <Typography
                                variant="button"
                            >

                                <div className="devname">นายธีรวัฒน์ กูดกิ่ง</div>
                            </Typography>
                        </Button>
                    </a>
                </div>
                <div className="column">
                    <a href="https://www.facebook.com/poln.jongketkam" className="fa">
                        <img src="https://media.tenor.com/LGCS8U0fTFkAAAAd/gojo-satoru-jjk.gif" alt="พล" width="270" height="260" className="avatar"></img>
                        <Button color="primary" variant="contained">
                            <Typography
                                variant="button"
                            >
                                <div className="devname">นายณัฐพล จงเกษกรรม</div>

                            </Typography>
                        </Button>
                    </a>
                </div>
                <div className="column">
                    <a href="https://www.facebook.com/gurocke.sus/" className="fa">
                        <img src="https://64.media.tumblr.com/ba41c54bc2342048c76ab1f04b7207d9/1cdf0f61bb8cf4e6-97/s500x750/61bde33934c020b9407225bdab9fad3b83fe3f2d.gif" alt="ฟิวส์" width="270" height="260" className="avatar"></img>
                        <Button color="primary" variant="contained">
                            <Typography
                                variant="button"
                            >
                                <div className="devname">นายอภิสิทธิ์ วงศ์วิศิษฐ์</div>

                            </Typography>
                        </Button>
                    </a>
                </div>
                <div className="column">
                    <a href="https://www.facebook.com/profile.php?id=100006438698029" className="fa">
                        <img src="https://media.tenor.com/mCQ06scNT8gAAAAS/flirt.gif" alt="มิว" width="270" height="260" className="avatar"></img>
                        <Button color="primary" variant="contained">
                            <Typography
                                variant="button"
                            >
                                <div className="devname">นายนายภานุพล จับชิ้น</div>
                            </Typography>
                        </Button>
                    </a>
                </div>
                <div className="column">
                    <a href="https://www.facebook.com/profile.php?id=100002751064253" className="fa">
                        <img src="https://cdn.discordapp.com/attachments/1049207840138604585/1073524666943164507/ezgif.com-crop_1.gif" alt="กัปตัน" width="270" height="260" className="avatar"></img>
                        <Button color="primary" variant="contained">
                            <Typography
                                variant="button"
                            >
                                <div className="devname">นายธนเดช เชิดในเมือง</div>
                            </Typography>
                        </Button>
                    </a>
                </div>
                <h1> <br /></h1>
            </div>

        </div>

    )



}
export default Homepage;