package main

// Algol68g - Algol 68 compiler-interpreter
// Homepage: https://jmvdveer.home.xs4all.nl/algol.html

import (
	"fmt"
	
	"os/exec"
)

func installAlgol68g() {
	// Método 1: Descargar y extraer .tar.gz
	algol68g_tar_url := "https://jmvdveer.home.xs4all.nl/algol68g-3.5.5.tar.gz"
	algol68g_cmd_tar := exec.Command("curl", "-L", algol68g_tar_url, "-o", "package.tar.gz")
	err := algol68g_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	algol68g_zip_url := "https://jmvdveer.home.xs4all.nl/algol68g-3.5.5.zip"
	algol68g_cmd_zip := exec.Command("curl", "-L", algol68g_zip_url, "-o", "package.zip")
	err = algol68g_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	algol68g_bin_url := "https://jmvdveer.home.xs4all.nl/algol68g-3.5.5.bin"
	algol68g_cmd_bin := exec.Command("curl", "-L", algol68g_bin_url, "-o", "binary.bin")
	err = algol68g_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	algol68g_src_url := "https://jmvdveer.home.xs4all.nl/algol68g-3.5.5.src.tar.gz"
	algol68g_cmd_src := exec.Command("curl", "-L", algol68g_src_url, "-o", "source.tar.gz")
	err = algol68g_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	algol68g_cmd_direct := exec.Command("./binary")
	err = algol68g_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}
