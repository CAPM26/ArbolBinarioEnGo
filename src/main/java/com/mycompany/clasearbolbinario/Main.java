/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.clasearbolbinario;

/**
 *
 * @author Usuario
 */
public class Main {
    
    public static void main(String[] args) {
        Arbol arbol = new Arbol();
        arbol.insertar(5);
        arbol.insertar(3);
        arbol.insertar(7);
        arbol.insertar(2);
        arbol.insertar(4);
        arbol.insertar(6);
        arbol.insertar(8);
//        System.out.println(arbol.existe(5));
//        System.out.println(arbol.existe(3));
//        System.out.println(arbol.existe(2));
//        System.out.println(arbol.existe(123455));
        arbol.dispararPreorden();
        System.out.println("\nfin de la Ejecucion" );
    
        }
}