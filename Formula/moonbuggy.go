package main

// MoonBuggy - Drive some car across the moon
// Homepage: https://www.seehuhn.de/pages/moon-buggy.html

import (
	"fmt"
	
	"os/exec"
)

func installMoonBuggy() {
	// Método 1: Descargar y extraer .tar.gz
	moonbuggy_tar_url := "https://m.seehuhn.de/programs/moon-buggy-1.0.tar.gz"
	moonbuggy_cmd_tar := exec.Command("curl", "-L", moonbuggy_tar_url, "-o", "package.tar.gz")
	err := moonbuggy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moonbuggy_zip_url := "https://m.seehuhn.de/programs/moon-buggy-1.0.zip"
	moonbuggy_cmd_zip := exec.Command("curl", "-L", moonbuggy_zip_url, "-o", "package.zip")
	err = moonbuggy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moonbuggy_bin_url := "https://m.seehuhn.de/programs/moon-buggy-1.0.bin"
	moonbuggy_cmd_bin := exec.Command("curl", "-L", moonbuggy_bin_url, "-o", "binary.bin")
	err = moonbuggy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moonbuggy_src_url := "https://m.seehuhn.de/programs/moon-buggy-1.0.src.tar.gz"
	moonbuggy_cmd_src := exec.Command("curl", "-L", moonbuggy_src_url, "-o", "source.tar.gz")
	err = moonbuggy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moonbuggy_cmd_direct := exec.Command("./binary")
	err = moonbuggy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
