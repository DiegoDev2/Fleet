package main

// Ripgrep - Search tool like grep and The Silver Searcher
// Homepage: https://github.com/BurntSushi/ripgrep

import (
	"fmt"
	
	"os/exec"
)

func installRipgrep() {
	// Método 1: Descargar y extraer .tar.gz
	ripgrep_tar_url := "https://github.com/BurntSushi/ripgrep/archive/refs/tags/14.1.0.tar.gz"
	ripgrep_cmd_tar := exec.Command("curl", "-L", ripgrep_tar_url, "-o", "package.tar.gz")
	err := ripgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ripgrep_zip_url := "https://github.com/BurntSushi/ripgrep/archive/refs/tags/14.1.0.zip"
	ripgrep_cmd_zip := exec.Command("curl", "-L", ripgrep_zip_url, "-o", "package.zip")
	err = ripgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ripgrep_bin_url := "https://github.com/BurntSushi/ripgrep/archive/refs/tags/14.1.0.bin"
	ripgrep_cmd_bin := exec.Command("curl", "-L", ripgrep_bin_url, "-o", "binary.bin")
	err = ripgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ripgrep_src_url := "https://github.com/BurntSushi/ripgrep/archive/refs/tags/14.1.0.src.tar.gz"
	ripgrep_cmd_src := exec.Command("curl", "-L", ripgrep_src_url, "-o", "source.tar.gz")
	err = ripgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ripgrep_cmd_direct := exec.Command("./binary")
	err = ripgrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
	exec.Command("latte", "install", "asciidoctor").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
