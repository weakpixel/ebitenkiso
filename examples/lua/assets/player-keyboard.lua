function draw()
    drawFilledCircle(get_x(), get_y(), 20)
end

function update()
    if isKeyPressed(PlayerUp) then
        move(0, -5)
    end
    if isKeyPressed(PlayerLeft) then
        move(-5, 0)
    end
    if isKeyPressed(PlayerRight) then
        move(5, 0)
    end
    if isKeyPressed(PlayerDown) then
        move(0, 5)
    end
end