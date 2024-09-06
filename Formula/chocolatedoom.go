package main

// ChocolateDoom - Accurate source port of Doom
// Homepage: https://www.chocolate-doom.org/

import (
	"fmt"
	
	"os/exec"
)

func installChocolateDoom() {
	// Método 1: Descargar y extraer .tar.gz
	chocolatedoom_tar_url := "https://github.com/chocolate-doom/chocolate-doom/archive/refs/tags/chocolate-doom-3.1.0.tar.gz"
	chocolatedoom_cmd_tar := exec.Command("curl", "-L", chocolatedoom_tar_url, "-o", "package.tar.gz")
	err := chocolatedoom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chocolatedoom_zip_url := "https://github.com/chocolate-doom/chocolate-doom/archive/refs/tags/chocolate-doom-3.1.0.zip"
	chocolatedoom_cmd_zip := exec.Command("curl", "-L", chocolatedoom_zip_url, "-o", "package.zip")
	err = chocolatedoom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chocolatedoom_bin_url := "https://github.com/chocolate-doom/chocolate-doom/archive/refs/tags/chocolate-doom-3.1.0.bin"
	chocolatedoom_cmd_bin := exec.Command("curl", "-L", chocolatedoom_bin_url, "-o", "binary.bin")
	err = chocolatedoom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chocolatedoom_src_url := "https://github.com/chocolate-doom/chocolate-doom/archive/refs/tags/chocolate-doom-3.1.0.src.tar.gz"
	chocolatedoom_cmd_src := exec.Command("curl", "-L", chocolatedoom_src_url, "-o", "source.tar.gz")
	err = chocolatedoom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chocolatedoom_cmd_direct := exec.Command("./binary")
	err = chocolatedoom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsamplerate")
	exec.Command("latte", "install", "libsamplerate").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: sdl2_net")
	exec.Command("latte", "install", "sdl2_net").Run()
}
