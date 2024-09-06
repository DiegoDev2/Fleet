package main

// Hexgui - GUI for playing Hex over Hex Text Protocol
// Homepage: https://sourceforge.net/p/benzene/hexgui/

import (
	"fmt"
	
	"os/exec"
)

func installHexgui() {
	// Método 1: Descargar y extraer .tar.gz
	hexgui_tar_url := "https://github.com/apetresc/hexgui/archive/refs/tags/v0.9.4.tar.gz"
	hexgui_cmd_tar := exec.Command("curl", "-L", hexgui_tar_url, "-o", "package.tar.gz")
	err := hexgui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hexgui_zip_url := "https://github.com/apetresc/hexgui/archive/refs/tags/v0.9.4.zip"
	hexgui_cmd_zip := exec.Command("curl", "-L", hexgui_zip_url, "-o", "package.zip")
	err = hexgui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hexgui_bin_url := "https://github.com/apetresc/hexgui/archive/refs/tags/v0.9.4.bin"
	hexgui_cmd_bin := exec.Command("curl", "-L", hexgui_bin_url, "-o", "binary.bin")
	err = hexgui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hexgui_src_url := "https://github.com/apetresc/hexgui/archive/refs/tags/v0.9.4.src.tar.gz"
	hexgui_cmd_src := exec.Command("curl", "-L", hexgui_src_url, "-o", "source.tar.gz")
	err = hexgui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hexgui_cmd_direct := exec.Command("./binary")
	err = hexgui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
	exec.Command("latte", "install", "ant").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
