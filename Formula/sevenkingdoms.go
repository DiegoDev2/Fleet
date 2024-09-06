package main

// SevenKingdoms - Real-time strategy game developed by Trevor Chan of Enlight Software
// Homepage: https://7kfans.com

import (
	"fmt"
	
	"os/exec"
)

func installSevenKingdoms() {
	// Método 1: Descargar y extraer .tar.gz
	sevenkingdoms_tar_url := "https://downloads.sourceforge.net/project/skfans/7KAA%202.15.6/7kaa-2.15.6.tar.gz"
	sevenkingdoms_cmd_tar := exec.Command("curl", "-L", sevenkingdoms_tar_url, "-o", "package.tar.gz")
	err := sevenkingdoms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sevenkingdoms_zip_url := "https://downloads.sourceforge.net/project/skfans/7KAA%202.15.6/7kaa-2.15.6.zip"
	sevenkingdoms_cmd_zip := exec.Command("curl", "-L", sevenkingdoms_zip_url, "-o", "package.zip")
	err = sevenkingdoms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sevenkingdoms_bin_url := "https://downloads.sourceforge.net/project/skfans/7KAA%202.15.6/7kaa-2.15.6.bin"
	sevenkingdoms_cmd_bin := exec.Command("curl", "-L", sevenkingdoms_bin_url, "-o", "binary.bin")
	err = sevenkingdoms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sevenkingdoms_src_url := "https://downloads.sourceforge.net/project/skfans/7KAA%202.15.6/7kaa-2.15.6.src.tar.gz"
	sevenkingdoms_cmd_src := exec.Command("curl", "-L", sevenkingdoms_src_url, "-o", "source.tar.gz")
	err = sevenkingdoms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sevenkingdoms_cmd_direct := exec.Command("./binary")
	err = sevenkingdoms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: enet")
	exec.Command("latte", "install", "enet").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openal-soft")
	exec.Command("latte", "install", "openal-soft").Run()
}
