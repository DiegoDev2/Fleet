package main

// Hivex - Library and tools for extracting the contents of Windows Registry hive files
// Homepage: https://libguestfs.org

import (
	"fmt"
	
	"os/exec"
)

func installHivex() {
	// Método 1: Descargar y extraer .tar.gz
	hivex_tar_url := "https://download.libguestfs.org/hivex/hivex-1.3.24.tar.gz"
	hivex_cmd_tar := exec.Command("curl", "-L", hivex_tar_url, "-o", "package.tar.gz")
	err := hivex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hivex_zip_url := "https://download.libguestfs.org/hivex/hivex-1.3.24.zip"
	hivex_cmd_zip := exec.Command("curl", "-L", hivex_zip_url, "-o", "package.zip")
	err = hivex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hivex_bin_url := "https://download.libguestfs.org/hivex/hivex-1.3.24.bin"
	hivex_cmd_bin := exec.Command("curl", "-L", hivex_bin_url, "-o", "binary.bin")
	err = hivex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hivex_src_url := "https://download.libguestfs.org/hivex/hivex-1.3.24.src.tar.gz"
	hivex_cmd_src := exec.Command("curl", "-L", hivex_src_url, "-o", "source.tar.gz")
	err = hivex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hivex_cmd_direct := exec.Command("./binary")
	err = hivex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
