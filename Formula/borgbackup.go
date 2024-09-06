package main

// Borgbackup - Deduplicating archiver with compression and authenticated encryption
// Homepage: https://borgbackup.org/

import (
	"fmt"
	
	"os/exec"
)

func installBorgbackup() {
	// Método 1: Descargar y extraer .tar.gz
	borgbackup_tar_url := "https://files.pythonhosted.org/packages/dd/0d/28e60180ce4ae171adba65ce9f8878fce3580c6d2cfdfa998929175105dd/borgbackup-1.4.0.tar.gz"
	borgbackup_cmd_tar := exec.Command("curl", "-L", borgbackup_tar_url, "-o", "package.tar.gz")
	err := borgbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	borgbackup_zip_url := "https://files.pythonhosted.org/packages/dd/0d/28e60180ce4ae171adba65ce9f8878fce3580c6d2cfdfa998929175105dd/borgbackup-1.4.0.zip"
	borgbackup_cmd_zip := exec.Command("curl", "-L", borgbackup_zip_url, "-o", "package.zip")
	err = borgbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	borgbackup_bin_url := "https://files.pythonhosted.org/packages/dd/0d/28e60180ce4ae171adba65ce9f8878fce3580c6d2cfdfa998929175105dd/borgbackup-1.4.0.bin"
	borgbackup_cmd_bin := exec.Command("curl", "-L", borgbackup_bin_url, "-o", "binary.bin")
	err = borgbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	borgbackup_src_url := "https://files.pythonhosted.org/packages/dd/0d/28e60180ce4ae171adba65ce9f8878fce3580c6d2cfdfa998929175105dd/borgbackup-1.4.0.src.tar.gz"
	borgbackup_cmd_src := exec.Command("curl", "-L", borgbackup_src_url, "-o", "source.tar.gz")
	err = borgbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	borgbackup_cmd_direct := exec.Command("./binary")
	err = borgbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libb2")
	exec.Command("latte", "install", "libb2").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: xxhash")
	exec.Command("latte", "install", "xxhash").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: acl")
	exec.Command("latte", "install", "acl").Run()
}
