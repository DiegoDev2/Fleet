package main

// Gbdfed - Bitmap Font Editor
// Homepage: http://sofia.nmsu.edu/~mleisher/Software/gbdfed/

import (
	"fmt"
	
	"os/exec"
)

func installGbdfed() {
	// Método 1: Descargar y extraer .tar.gz
	gbdfed_tar_url := "http://sofia.nmsu.edu/~mleisher/Software/gbdfed/gbdfed-1.6.tar.gz"
	gbdfed_cmd_tar := exec.Command("curl", "-L", gbdfed_tar_url, "-o", "package.tar.gz")
	err := gbdfed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gbdfed_zip_url := "http://sofia.nmsu.edu/~mleisher/Software/gbdfed/gbdfed-1.6.zip"
	gbdfed_cmd_zip := exec.Command("curl", "-L", gbdfed_zip_url, "-o", "package.zip")
	err = gbdfed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gbdfed_bin_url := "http://sofia.nmsu.edu/~mleisher/Software/gbdfed/gbdfed-1.6.bin"
	gbdfed_cmd_bin := exec.Command("curl", "-L", gbdfed_bin_url, "-o", "binary.bin")
	err = gbdfed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gbdfed_src_url := "http://sofia.nmsu.edu/~mleisher/Software/gbdfed/gbdfed-1.6.src.tar.gz"
	gbdfed_cmd_src := exec.Command("curl", "-L", gbdfed_src_url, "-o", "source.tar.gz")
	err = gbdfed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gbdfed_cmd_direct := exec.Command("./binary")
	err = gbdfed_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
}
