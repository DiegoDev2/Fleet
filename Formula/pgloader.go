package main

// Pgloader - Data loading tool for PostgreSQL
// Homepage: https://github.com/dimitri/pgloader

import (
	"fmt"
	
	"os/exec"
)

func installPgloader() {
	// Método 1: Descargar y extraer .tar.gz
	pgloader_tar_url := "https://github.com/dimitri/pgloader/releases/download/v3.6.9/pgloader-bundle-3.6.9.tgz"
	pgloader_cmd_tar := exec.Command("curl", "-L", pgloader_tar_url, "-o", "package.tar.gz")
	err := pgloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgloader_zip_url := "https://github.com/dimitri/pgloader/releases/download/v3.6.9/pgloader-bundle-3.6.9.tgz"
	pgloader_cmd_zip := exec.Command("curl", "-L", pgloader_zip_url, "-o", "package.zip")
	err = pgloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgloader_bin_url := "https://github.com/dimitri/pgloader/releases/download/v3.6.9/pgloader-bundle-3.6.9.tgz"
	pgloader_cmd_bin := exec.Command("curl", "-L", pgloader_bin_url, "-o", "binary.bin")
	err = pgloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgloader_src_url := "https://github.com/dimitri/pgloader/releases/download/v3.6.9/pgloader-bundle-3.6.9.tgz"
	pgloader_cmd_src := exec.Command("curl", "-L", pgloader_src_url, "-o", "source.tar.gz")
	err = pgloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgloader_cmd_direct := exec.Command("./binary")
	err = pgloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: buildapp")
	exec.Command("latte", "install", "buildapp").Run()
	fmt.Println("Instalando dependencia: freetds")
	exec.Command("latte", "install", "freetds").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: sbcl")
	exec.Command("latte", "install", "sbcl").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
