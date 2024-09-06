package main

// Aide - File and directory integrity checker
// Homepage: https://aide.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installAide() {
	// Método 1: Descargar y extraer .tar.gz
	aide_tar_url := "https://github.com/aide/aide/releases/download/v0.18.8/aide-0.18.8.tar.gz"
	aide_cmd_tar := exec.Command("curl", "-L", aide_tar_url, "-o", "package.tar.gz")
	err := aide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aide_zip_url := "https://github.com/aide/aide/releases/download/v0.18.8/aide-0.18.8.zip"
	aide_cmd_zip := exec.Command("curl", "-L", aide_zip_url, "-o", "package.zip")
	err = aide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aide_bin_url := "https://github.com/aide/aide/releases/download/v0.18.8/aide-0.18.8.bin"
	aide_cmd_bin := exec.Command("curl", "-L", aide_bin_url, "-o", "binary.bin")
	err = aide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aide_src_url := "https://github.com/aide/aide/releases/download/v0.18.8/aide-0.18.8.src.tar.gz"
	aide_cmd_src := exec.Command("curl", "-L", aide_src_url, "-o", "source.tar.gz")
	err = aide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aide_cmd_direct := exec.Command("./binary")
	err = aide_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: autoconf-archive")
	exec.Command("latte", "install", "autoconf-archive").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
