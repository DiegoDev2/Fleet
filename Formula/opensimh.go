package main

// OpenSimh - Multi-system computer simulator
// Homepage: https://opensimh.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenSimh() {
	// Método 1: Descargar y extraer .tar.gz
	opensimh_tar_url := "https://github.com/open-simh/simh/archive/refs/tags/v3.12-3.tar.gz"
	opensimh_cmd_tar := exec.Command("curl", "-L", opensimh_tar_url, "-o", "package.tar.gz")
	err := opensimh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opensimh_zip_url := "https://github.com/open-simh/simh/archive/refs/tags/v3.12-3.zip"
	opensimh_cmd_zip := exec.Command("curl", "-L", opensimh_zip_url, "-o", "package.zip")
	err = opensimh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opensimh_bin_url := "https://github.com/open-simh/simh/archive/refs/tags/v3.12-3.bin"
	opensimh_cmd_bin := exec.Command("curl", "-L", opensimh_bin_url, "-o", "binary.bin")
	err = opensimh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opensimh_src_url := "https://github.com/open-simh/simh/archive/refs/tags/v3.12-3.src.tar.gz"
	opensimh_cmd_src := exec.Command("curl", "-L", opensimh_src_url, "-o", "source.tar.gz")
	err = opensimh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opensimh_cmd_direct := exec.Command("./binary")
	err = opensimh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: vde")
	exec.Command("latte", "install", "vde").Run()
}
