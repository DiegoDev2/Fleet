package main

// Plowshare - Download/upload tool for popular file sharing websites
// Homepage: https://github.com/mcrapet/plowshare

import (
	"fmt"
	
	"os/exec"
)

func installPlowshare() {
	// Método 1: Descargar y extraer .tar.gz
	plowshare_tar_url := "https://github.com/mcrapet/plowshare/archive/refs/tags/v2.1.7.tar.gz"
	plowshare_cmd_tar := exec.Command("curl", "-L", plowshare_tar_url, "-o", "package.tar.gz")
	err := plowshare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plowshare_zip_url := "https://github.com/mcrapet/plowshare/archive/refs/tags/v2.1.7.zip"
	plowshare_cmd_zip := exec.Command("curl", "-L", plowshare_zip_url, "-o", "package.zip")
	err = plowshare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plowshare_bin_url := "https://github.com/mcrapet/plowshare/archive/refs/tags/v2.1.7.bin"
	plowshare_cmd_bin := exec.Command("curl", "-L", plowshare_bin_url, "-o", "binary.bin")
	err = plowshare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plowshare_src_url := "https://github.com/mcrapet/plowshare/archive/refs/tags/v2.1.7.src.tar.gz"
	plowshare_cmd_src := exec.Command("curl", "-L", plowshare_src_url, "-o", "source.tar.gz")
	err = plowshare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plowshare_cmd_direct := exec.Command("./binary")
	err = plowshare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
exec.Command("latte", "install", "bash")
	fmt.Println("Instalando dependencia: feh")
exec.Command("latte", "install", "feh")
	fmt.Println("Instalando dependencia: libcaca")
exec.Command("latte", "install", "libcaca")
	fmt.Println("Instalando dependencia: recode")
exec.Command("latte", "install", "recode")
	fmt.Println("Instalando dependencia: spidermonkey")
exec.Command("latte", "install", "spidermonkey")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
}
