package main

// Takt - Text-based music programming language
// Homepage: https://takt.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installTakt() {
	// Método 1: Descargar y extraer .tar.gz
	takt_tar_url := "https://downloads.sourceforge.net/project/takt/takt-0.310-src.tar.gz"
	takt_cmd_tar := exec.Command("curl", "-L", takt_tar_url, "-o", "package.tar.gz")
	err := takt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	takt_zip_url := "https://downloads.sourceforge.net/project/takt/takt-0.310-src.zip"
	takt_cmd_zip := exec.Command("curl", "-L", takt_zip_url, "-o", "package.zip")
	err = takt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	takt_bin_url := "https://downloads.sourceforge.net/project/takt/takt-0.310-src.bin"
	takt_cmd_bin := exec.Command("curl", "-L", takt_bin_url, "-o", "binary.bin")
	err = takt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	takt_src_url := "https://downloads.sourceforge.net/project/takt/takt-0.310-src.src.tar.gz"
	takt_cmd_src := exec.Command("curl", "-L", takt_src_url, "-o", "source.tar.gz")
	err = takt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	takt_cmd_direct := exec.Command("./binary")
	err = takt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
