package main

// Icemon - Icecream GUI Monitor
// Homepage: https://github.com/icecc/icemon

import (
	"fmt"
	
	"os/exec"
)

func installIcemon() {
	// Método 1: Descargar y extraer .tar.gz
	icemon_tar_url := "https://github.com/icecc/icemon/archive/refs/tags/v3.3.tar.gz"
	icemon_cmd_tar := exec.Command("curl", "-L", icemon_tar_url, "-o", "package.tar.gz")
	err := icemon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icemon_zip_url := "https://github.com/icecc/icemon/archive/refs/tags/v3.3.zip"
	icemon_cmd_zip := exec.Command("curl", "-L", icemon_zip_url, "-o", "package.zip")
	err = icemon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icemon_bin_url := "https://github.com/icecc/icemon/archive/refs/tags/v3.3.bin"
	icemon_cmd_bin := exec.Command("curl", "-L", icemon_bin_url, "-o", "binary.bin")
	err = icemon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icemon_src_url := "https://github.com/icecc/icemon/archive/refs/tags/v3.3.src.tar.gz"
	icemon_cmd_src := exec.Command("curl", "-L", icemon_src_url, "-o", "source.tar.gz")
	err = icemon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icemon_cmd_direct := exec.Command("./binary")
	err = icemon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: extra-cmake-modules")
	exec.Command("latte", "install", "extra-cmake-modules").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: icecream")
	exec.Command("latte", "install", "icecream").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libcap-ng")
	exec.Command("latte", "install", "libcap-ng").Run()
}
