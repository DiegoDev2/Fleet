package main

// Orbuculum - Arm Cortex-M SWO/SWV Demux and Postprocess
// Homepage: https://github.com/orbcode/orbuculum

import (
	"fmt"
	
	"os/exec"
)

func installOrbuculum() {
	// Método 1: Descargar y extraer .tar.gz
	orbuculum_tar_url := "https://github.com/orbcode/orbuculum/archive/refs/tags/V2.1.0.tar.gz"
	orbuculum_cmd_tar := exec.Command("curl", "-L", orbuculum_tar_url, "-o", "package.tar.gz")
	err := orbuculum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orbuculum_zip_url := "https://github.com/orbcode/orbuculum/archive/refs/tags/V2.1.0.zip"
	orbuculum_cmd_zip := exec.Command("curl", "-L", orbuculum_zip_url, "-o", "package.zip")
	err = orbuculum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orbuculum_bin_url := "https://github.com/orbcode/orbuculum/archive/refs/tags/V2.1.0.bin"
	orbuculum_cmd_bin := exec.Command("curl", "-L", orbuculum_bin_url, "-o", "binary.bin")
	err = orbuculum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orbuculum_src_url := "https://github.com/orbcode/orbuculum/archive/refs/tags/V2.1.0.src.tar.gz"
	orbuculum_cmd_src := exec.Command("curl", "-L", orbuculum_src_url, "-o", "source.tar.gz")
	err = orbuculum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orbuculum_cmd_direct := exec.Command("./binary")
	err = orbuculum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}
