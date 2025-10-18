function draw()
    drawFilledCircle(get_x(), get_y(), 20)
end

function update()
    mx, my = cursorPosition()
    set_x(mx)
    set_y(my)

    if isMouseButtonPressed(LeftButton) then
        set_data("clicked", true)
    end 
    
end