package main

// MagicWormholeRs - Rust implementation of Magic Wormhole, with new features and enhancements
// Homepage: https://github.com/magic-wormhole/magic-wormhole.rs

import (
	"fmt"
	
	"os/exec"
)

func installMagicWormholeRs() {
	// Método 1: Descargar y extraer .tar.gz
	magicwormholers_tar_url := "https://github.com/magic-wormhole/magic-wormhole.rs/archive/refs/tags/0.7.1.tar.gz"
	magicwormholers_cmd_tar := exec.Command("curl", "-L", magicwormholers_tar_url, "-o", "package.tar.gz")
	err := magicwormholers_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	magicwormholers_zip_url := "https://github.com/magic-wormhole/magic-wormhole.rs/archive/refs/tags/0.7.1.zip"
	magicwormholers_cmd_zip := exec.Command("curl", "-L", magicwormholers_zip_url, "-o", "package.zip")
	err = magicwormholers_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	magicwormholers_bin_url := "https://github.com/magic-wormhole/magic-wormhole.rs/archive/refs/tags/0.7.1.bin"
	magicwormholers_cmd_bin := exec.Command("curl", "-L", magicwormholers_bin_url, "-o", "binary.bin")
	err = magicwormholers_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	magicwormholers_src_url := "https://github.com/magic-wormhole/magic-wormhole.rs/archive/refs/tags/0.7.1.src.tar.gz"
	magicwormholers_cmd_src := exec.Command("curl", "-L", magicwormholers_src_url, "-o", "source.tar.gz")
	err = magicwormholers_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	magicwormholers_cmd_direct := exec.Command("./binary")
	err = magicwormholers_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
