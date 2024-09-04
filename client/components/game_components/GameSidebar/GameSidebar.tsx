import "./GameSidebar.scss";

export default function GameSidebar() {
  return (
    <aside className="game-sidebar">
      <div className="game-sidebar-currency">
        <div style={{ display: "flex" }}>
          <div className="test-yellow">Gold"icon"</div>: 842 678K
        </div>
        <div style={{ display: "flex" }}>
          <div className="test-blue">Diamond"icon"</div>: 450
        </div>
      </div>
      <div className="game-sidebar-options">
        <div className="game-sidebar-tabs">
          <div className="game-sidebar-tab-el">One</div>
          <div className="game-sidebar-tab-el">Two</div>
          <div className="game-sidebar-tab-el">Three</div>
          <div className="game-sidebar-tab-el">Four</div>
        </div>
        <div className="game-sidebar-options-content">
          {" "}
          {/*Wrapper */}
          <div className="game-sidebar-options-content-title">Tab title</div>
          <div className="game-sidebar-options-content-scroll">
            <div className="scroll-element">
              <div className="scroll-element-info-wrapper">
                <div className="scroll-element-title">Element 1 title</div>
                <div className="scroll-element-description">
                  Brief element description
                </div>
              </div>
              <div>Action buttons?</div>
            </div>
            <div className="scroll-element">
              <div className="scroll-element-info-wrapper">
                <div className="scroll-element-title">Element 2 title</div>
                <div className="scroll-element-description">
                  Brief element description
                </div>
              </div>
              <div>Action buttons?</div>
            </div>
          </div>
        </div>
      </div>
    </aside>
  );
}
