package main

// Alda - Music programming language for musicians
// Homepage: https://alda.io

import (
	"fmt"
	
	"os/exec"
)

func installAlda() {
	// Método 1: Descargar y extraer .tar.gz
	alda_tar_url := "https://github.com/alda-lang/alda/archive/refs/tags/release-2.3.1.tar.gz"
	alda_cmd_tar := exec.Command("curl", "-L", alda_tar_url, "-o", "package.tar.gz")
	err := alda_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alda_zip_url := "https://github.com/alda-lang/alda/archive/refs/tags/release-2.3.1.zip"
	alda_cmd_zip := exec.Command("curl", "-L", alda_zip_url, "-o", "package.zip")
	err = alda_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alda_bin_url := "https://github.com/alda-lang/alda/archive/refs/tags/release-2.3.1.bin"
	alda_cmd_bin := exec.Command("curl", "-L", alda_bin_url, "-o", "binary.bin")
	err = alda_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alda_src_url := "https://github.com/alda-lang/alda/archive/refs/tags/release-2.3.1.src.tar.gz"
	alda_cmd_src := exec.Command("curl", "-L", alda_src_url, "-o", "source.tar.gz")
	err = alda_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alda_cmd_direct := exec.Command("./binary")
	err = alda_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: gradle")
	exec.Command("latte", "install", "gradle").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
