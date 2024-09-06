package main

// TranslateShell - Command-line translator using Google Translate and more
// Homepage: https://www.soimort.org/translate-shell

import (
	"fmt"
	
	"os/exec"
)

func installTranslateShell() {
	// Método 1: Descargar y extraer .tar.gz
	translateshell_tar_url := "https://github.com/soimort/translate-shell/archive/refs/tags/v0.9.7.1.tar.gz"
	translateshell_cmd_tar := exec.Command("curl", "-L", translateshell_tar_url, "-o", "package.tar.gz")
	err := translateshell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	translateshell_zip_url := "https://github.com/soimort/translate-shell/archive/refs/tags/v0.9.7.1.zip"
	translateshell_cmd_zip := exec.Command("curl", "-L", translateshell_zip_url, "-o", "package.zip")
	err = translateshell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	translateshell_bin_url := "https://github.com/soimort/translate-shell/archive/refs/tags/v0.9.7.1.bin"
	translateshell_cmd_bin := exec.Command("curl", "-L", translateshell_bin_url, "-o", "binary.bin")
	err = translateshell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	translateshell_src_url := "https://github.com/soimort/translate-shell/archive/refs/tags/v0.9.7.1.src.tar.gz"
	translateshell_cmd_src := exec.Command("curl", "-L", translateshell_src_url, "-o", "source.tar.gz")
	err = translateshell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	translateshell_cmd_direct := exec.Command("./binary")
	err = translateshell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fribidi")
	exec.Command("latte", "install", "fribidi").Run()
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: rlwrap")
	exec.Command("latte", "install", "rlwrap").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
