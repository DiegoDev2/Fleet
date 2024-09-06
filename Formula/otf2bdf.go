package main

// Otf2bdf - OpenType to BDF font converter
// Homepage: https://github.com/jirutka/otf2bdf

import (
	"fmt"
	
	"os/exec"
)

func installOtf2bdf() {
	// Método 1: Descargar y extraer .tar.gz
	otf2bdf_tar_url := "https://slackware.uk/~urchlay/src/otf2bdf-3.1.tbz2"
	otf2bdf_cmd_tar := exec.Command("curl", "-L", otf2bdf_tar_url, "-o", "package.tar.gz")
	err := otf2bdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	otf2bdf_zip_url := "https://slackware.uk/~urchlay/src/otf2bdf-3.1.tbz2"
	otf2bdf_cmd_zip := exec.Command("curl", "-L", otf2bdf_zip_url, "-o", "package.zip")
	err = otf2bdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	otf2bdf_bin_url := "https://slackware.uk/~urchlay/src/otf2bdf-3.1.tbz2"
	otf2bdf_cmd_bin := exec.Command("curl", "-L", otf2bdf_bin_url, "-o", "binary.bin")
	err = otf2bdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	otf2bdf_src_url := "https://slackware.uk/~urchlay/src/otf2bdf-3.1.tbz2"
	otf2bdf_cmd_src := exec.Command("curl", "-L", otf2bdf_src_url, "-o", "source.tar.gz")
	err = otf2bdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	otf2bdf_cmd_direct := exec.Command("./binary")
	err = otf2bdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
}
