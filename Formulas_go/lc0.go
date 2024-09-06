package main

// Lc0 - Open source neural network based chess engine
// Homepage: https://lczero.org/

import (
	"fmt"
	
	"os/exec"
)

func installLc0() {
	// Método 1: Descargar y extraer .tar.gz
	lc0_tar_url := "https://github.com/LeelaChessZero/lc0.git"
	lc0_cmd_tar := exec.Command("curl", "-L", lc0_tar_url, "-o", "package.tar.gz")
	err := lc0_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lc0_zip_url := "https://github.com/LeelaChessZero/lc0.git"
	lc0_cmd_zip := exec.Command("curl", "-L", lc0_zip_url, "-o", "package.zip")
	err = lc0_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lc0_bin_url := "https://github.com/LeelaChessZero/lc0.git"
	lc0_cmd_bin := exec.Command("curl", "-L", lc0_bin_url, "-o", "binary.bin")
	err = lc0_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lc0_src_url := "https://github.com/LeelaChessZero/lc0.git"
	lc0_cmd_src := exec.Command("curl", "-L", lc0_src_url, "-o", "source.tar.gz")
	err = lc0_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lc0_cmd_direct := exec.Command("./binary")
	err = lc0_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
}
